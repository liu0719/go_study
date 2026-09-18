package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ============================================================
// 第 0 步：Cookie 到底是什么（先把这个理解透，后面三种方案全是它的变体）
//
//   Cookie 不是"登录态"，它只是一个【浏览器替你自动携带的文本容器】。
//
//   服务端 → 响应头：Set-Cookie: sid=abc; Path=/; HttpOnly
//   浏览器 → 之后每次请求头：Cookie: sid=abc
//
//   所以"用 Cookie 做登录"的完整语义是：Cookie 里放一个【凭据】。
//   这个凭据要么自己能被验证（方案 C，无状态），
//   要么服务端拿它去查表（方案 B，有状态）。
// ============================================================

// 用环境变量区分实例，方便起两个进程复现"多实例丢登录态"
var (
	port       string
	instanceID string
)

// ---------- 方案 A：只玩 Cookie 本身，先把 Set-Cookie 跑通 ----------

func setCookieDemo(c *gin.Context) {
	// gin 的 SetCookie 里 maxage 单位是【秒】；传 0 表示不设过期时间（随浏览器关闭失效）
	// 参数顺序：name, value, maxage, path, domain, secure, httponly
	c.SetCookie("uid", "u_88888", int(24*time.Hour/time.Second), "/", "localhost", false, true)
}

func getCookieDemo(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录：Cookie 不存在", "hint": "先访问 /cookie/set"})
		return
	}
	// 打印真实报文，亲眼看一次"浏览器真的把它带回来了"
	fmt.Printf("[收到 Cookie 报文] %v\n", c.Request.Header.Values("Cookie"))
	c.JSON(200, gin.H{"msg": "已登录", "uid": uid})
}

// ---------- 方案 B：进程内 Session（有状态：真数据存在服务端）----------
//
// 关键区别：Cookie 里只放一个 sessionId，真正的登录数据在服务器内存里。
// 好处是随时可以踢人下线（删掉表里的记录即可），坏处是服务器必须有状态。

type Session struct {
	UserID   string
	LoginAt  time.Time
	ExpireAt time.Time
}

type MemorySessionStore struct {
	mu   sync.RWMutex
	data map[string]*Session // key = sessionId（存在 Cookie 里） value = 真实登录态
}

func NewMemorySessionStore() *MemorySessionStore {
	return &MemorySessionStore{data: make(map[string]*Session)}
}

func (s *MemorySessionStore) Get(id string) (*Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	se, ok := s.data[id]
	if !ok {
		return nil, false
	}
	// 惰性过期检查：读的时候顺便判断是否过期
	if time.Now().After(se.ExpireAt) {
		return nil, false
	}
	return se, true
}

func (s *MemorySessionStore) Set(id string, se *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[id] = se
}

func (s *MemorySessionStore) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, id)
}

func (s *MemorySessionStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// 生成一个随机 sessionId，模拟真实框架给你的那个随机串
func newID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

// session 中间件：进来先找 Cookie 里的 sid，找不到就发一个。
// 它同时把 store 塞进 Context，所以下面的 handler 才能 c.MustGet("store") 拿到。
func sessionMiddleware(store *MemorySessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("store", store)

		sid, err := c.Cookie("sid")
		_, ok := store.Get(sid)
		if err != nil || !ok {
			// 两种情况：从没登录过 / 登录过但 session 过期或被别的实例创建。
			// 这里为了演示方便直接新建；真实项目应改为返回 401 让前端跳登录页。
			sid = newID()
			store.Set(sid, &Session{
				UserID:   "游客",
				LoginAt:  time.Now(),
				ExpireAt: time.Now().Add(30 * time.Minute),
			})
			c.SetCookie("sid", sid, int(30*time.Minute/time.Second), "/", "localhost", false, true)
		}
		c.Set("sid", sid)
		c.Next()
	}
}

func sessionLogin(c *gin.Context) {
	store := c.MustGet("store").(*MemorySessionStore)
	sid := c.MustGet("sid").(string)
	store.Set(sid, &Session{
		UserID:   "张三",
		LoginAt:  time.Now(),
		ExpireAt: time.Now().Add(30 * time.Minute),
	})
	// 注意：sessionId 不变，只是把表里的数据换了 → 这就是"有状态"的代价，
	// 你必须能准确定位到【那台持有这份数据的机器】
	c.JSON(200, gin.H{"msg": "登录成功", "sessionId": sid})
}

func sessionMe(c *gin.Context) {
	store := c.MustGet("store").(*MemorySessionStore)
	se, ok := store.Get(c.MustGet("sid").(string))
	if !ok || se.UserID == "游客" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录"})
		return
	}
	c.JSON(200, gin.H{
		"msg":    "已登录",
		"userID": se.UserID,
		"实例":     instanceID,
		"本机在线数":  store.Count(),
		"到期时间":   se.ExpireAt.Format(time.RFC3339),
	})
}

func sessionLogout(c *gin.Context) {
	store := c.MustGet("store").(*MemorySessionStore)
	sid := c.MustGet("sid").(string)
	store.Delete(sid) // ← 一句话踢人下线，这是 Session 最大的好处
	// 把 Cookie 的过期时间设到过去，让浏览器删掉它
	c.SetCookie("sid", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"msg": "已退出，本机在线数=" + fmt.Sprint(store.Count())})
}

// ---------- 方案 C：无状态签名 Token（JWT 的核心思想，纯标准库实现）----------
//
// 结构：base64(header).base64(payload).hmac签名
// 价值：服务端不存任何东西，任何人拿到 token 都能验证 → 天然适配多实例、水平扩容。
// 代价：服务端无法主动失效 token（除非维护黑名单，那就又变回有状态了）

const jwtSecret = "change-me-in-production"

func hmacSHA256(msg, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msg))
	return h.Sum(nil)
}

func signToken(payload map[string]any) (string, error) {
	header, _ := json.Marshal(map[string]string{"alg": "HS256", "typ": "JWT"})
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	h := base64.RawURLEncoding.EncodeToString(header)
	p := base64.RawURLEncoding.EncodeToString(body)
	signing := h + "." + p // 只对前两段做签名，签名放最后
	return signing + "." + base64.RawURLEncoding.EncodeToString(hmacSHA256(signing, jwtSecret)), nil
}

// 验签 + 过期检查，两段都要过才算合法
func verifyToken(token string) (map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("格式错误：应该是 3 段，实际 %d 段", len(parts))
	}
	expect := base64.RawURLEncoding.EncodeToString(hmacSHA256(parts[0]+"."+parts[1], jwtSecret))
	if !hmac.Equal([]byte(expect), []byte(parts[2])) {
		return nil, fmt.Errorf("签名不匹配：payload 被篡改过")
	}
	var payload map[string]any
	dec, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("payload 解码失败")
	}
	if err := json.Unmarshal(dec, &payload); err != nil {
		return nil, fmt.Errorf("payload 不是合法 JSON")
	}
	if exp, ok := payload["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
		return nil, fmt.Errorf("token 已过期")
	}
	return payload, nil
}

func jwtLogin(c *gin.Context) {
	token, err := signToken(map[string]any{
		"sub":  "zhangsan", // subject，用户标识
		"role": "user",
		"exp":  time.Now().Add(time.Hour).Unix(), // 过期时间，Unix 秒
	})
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	// 同时放 Cookie 和 Header 两种传法，让你体会区别
	// 放 Cookie 的优点：前端不用手动带，浏览器自动带；缺点：必须处理 CSRF
	// 放 Header 的优点：天然跨域友好、不受 SameSite 限制；缺点：JS 里要自己存（建议 sessionStorage，别用 localStorage，XSS 了就被抓走）
	c.SetCookie("jwt", token, int(time.Hour/time.Second), "/", "localhost", false, false) // 注意 httponly=false，演示时方便前端读出来给你看
	c.JSON(200, gin.H{
		"msg":   "登录成功",
		"token": token,
		"提示":    "第一段第二段用 base64 编码而不是加密，所以 payload 内容一眼就能看懂 —— JWT 防篡改，不防偷看",
	})
}

func jwtMe(c *gin.Context) {
	// 真实项目一般是 Authorization: Bearer <token>，这里 Cookie 和 Header 都支持
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		token, _ = c.Cookie("jwt")
	}
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录：没有 token", "hint": "先访问 /jwt/login"})
		return
	}
	payload, err := verifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "token 无效", "原因": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "已登录", "payload": payload, "实例": instanceID})
}

// 考点演示：篡改 payload 后再验签，看它怎么被识破
func jwtTamper(c *gin.Context) {
	good, _ := signToken(map[string]any{"sub": "zhangsan", "role": "user", "exp": time.Now().Add(time.Hour).Unix()})
	// 把 payload 里的 role 从 user 改成 admin，签名不动
	payload, _ := json.Marshal(map[string]any{"sub": "zhangsan", "role": "admin", "exp": time.Now().Add(time.Hour).Unix()})
	forged := good[:strings.Index(good, ".")] + "." + base64.RawURLEncoding.EncodeToString(payload) + good[strings.LastIndex(good, "."):]

	goodRes, goodErr := verifyToken(good)
	forgedRes, forgedErr := verifyToken(forged)

	c.JSON(200, gin.H{
		"合法token":   good,
		"合法token验签": fmt.Sprintf("结果=%v err=%v", goodRes != nil, goodErr),
		"伪造token":   forged,
		"伪造token验签": fmt.Sprintf("结果=%v err=%v", forgedRes != nil, forgedErr),
		"结论":        "改了 payload 就换不了签名，因为签名是对 payload 算出来的。但注意：如果 jwtSecret 泄露，攻击者可以自己签任意 token",
	})
}

// ---------- 第 5 步：全篇最重要 —— 亲眼复现"Session 在多实例下丢登录态" ----------
//
// 做法：开两个终端起两个实例，A 登录拿 sid，把 sid 拿去访问 B。
// 结果：B 说未登录。因为 B 的内存里从来没有过这个 sid。
// 这就是"生产上 Session 必须挂 Redis 共享存储"的一手证据。

func about(c *gin.Context) {
	store := c.MustGet("store").(*MemorySessionStore)
	c.JSON(200, gin.H{
		"实例ID":  instanceID,
		"本机会话数": store.Count(),
		"端口":    port,
		"三种方案对比": gin.H{
			"方案B_Session": "服务端存真数据。能随时踢人下线，但多实例必须共享存储(Redis)，否则登录态在负载均衡下随机丢失",
			"方案C_JWT":     "服务端不存。天然支持多实例和扩容，但无法主动失效 token",
			"方案A_裸Cookie": "只是演示 Cookie 机制，没有任何安全性，生产不要用",
		},
	})
}

func main() {
	// 环境变量控制，方便起两个实例做对比实验
	port = os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	instanceID = os.Getenv("INSTANCE_ID")
	if instanceID == "" {
		instanceID = "A"
	}

	r := gin.Default()

	store := NewMemorySessionStore()
	sessionMW := sessionMiddleware(store) // 中间件可以像 handler 一样按路由单独挂

	r.LoadHTMLFiles("16.cookie和session处理/登录测试.html")
	r.GET("", func(c *gin.Context) { c.HTML(200, "登录测试.html", "") })

	// 方案 A：裸 Cookie
	r.GET("/cookie/set", setCookieDemo)
	r.GET("/cookie/get", getCookieDemo)

	// 方案 B：Session（每个路由都要挂 sessionMW，因为它负责发/续 sid）
	r.GET("/session/login", sessionMW, sessionLogin)
	r.GET("/session/me", sessionMW, sessionMe)
	r.GET("/session/logout", sessionMW, sessionLogout)

	// 方案 C：无状态签名 Token
	r.GET("/jwt/login", jwtLogin)
	r.GET("/jwt/me", jwtMe)
	r.GET("/jwt/tamper", jwtTamper)

	r.GET("/about", sessionMW, about)

	fmt.Printf("======== Cookie / Session 演示服务 ========\n实例ID=%s  监听=%s\n浏览器打开 http://localhost%s 或直接看 /about\n", instanceID, port, port)
	r.Run(port)
}
