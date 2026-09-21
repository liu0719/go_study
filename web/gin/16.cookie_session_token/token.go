package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//
// 一句话总纲：
// 任何人（包括负载均衡后面那台陌生机器）只要拿同一把钥匙重算一次签名，就能判断 token 有没有被改过。
// 服务端因此一个字节都不用存——这就是"无状态"的来源。
//
// 结构：base64(header).base64(payload).hmac签名
// 价值：服务端不存任何东西，任何人拿到 token 都能验证 → 天然适配多实例、水平扩容。
// 代价：服务端无法主动失效 token（除非维护黑名单，那就又变回有状态了）
//
// 【过一遍的底线】下面几十行你只需要记住 4 条，其余全是实现细节：
//   1. 前两段是 base64 【编码】不是【加密】 → 客户端能直接解出内容（防篡改，不防偷看）
//   2. 签名是拿【前两段拼起来的字符串】算出来的 → 改 payload 一段，签名立刻对不上
//   3. 验签通过 ≠ 一定合法，还要单独查 exp 过期时间（签名只保证"没被改"，不保证"没过期"）
//   4. 整条链路的安全性只压在 jwtSecret 不能泄露这一件事上
//
// 【对应 golang-jwt/jwt/v5】等你直接调库时，下面这段手写代码可以整体忘掉，映射关系是：
//   signToken   →  jwt.NewWithClaims(method, claims) + token.SignedString(key)
//   verifyToken →  jwt.ParseWithClaims(str, claims, keyfunc)
//   hmacSHA256  →  库内部的 m.Sign()，黑盒，你不用再碰

const jwtSecret = "change-me-in-production"

// ↑ 真实项目绝对不要写死在代码里：git 提交即泄露，等于把签名权公开给全世界。
//   生产做法是从环境变量 / 配置中心读，且用足够长的随机串（≥32 字节）。
//   这里写常量只为演示，值故意写成一个"看到就该报警"的样式。

// hmacSHA256：把"消息 + 密钥"变成一个定长指纹。
// 内部结构（SHA-256 怎么算的）不用学，只需要记住两条性质，够用了：
//   - 同一条消息 + 同一个密钥 → 结果永远相同（可复现，这是"验签"能成立的前提）
//   - 消息改一个字节 → 结果完全不同（防篡改的依据）
//
// 注意它是哈希不是加密，不能反推原文；这里全程当黑盒用。
func hmacSHA256(msg, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret)) // 先创建"带着这把钥匙"的哈希器
	h.Write([]byte(msg))                      // 再喂入要签名的内容
	return h.Sum(nil)                         // 取出最终指纹
}

// signToken：把 payload 打成"三段式字符串"。对应 golang-jwt 的 NewWithClaims + SignedString。
func signToken(payload map[string]any) (string, error) {
	// 第一段 header：声明"我用的什么算法"。固定写法，实际项目中几乎不会改
	header, _ := json.Marshal(map[string]string{"alg": "HS256", "typ": "JWT"})
	// 第二段 payload：真正装业务数据的地方（谁登录的、什么权限、何时过期）
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	// base64 的作用只是"把 JSON 变成可以安全拼进字符串的字符"，任何人都能反向解回来
	h := base64.RawURLEncoding.EncodeToString(header)
	p := base64.RawURLEncoding.EncodeToString(body)
	// ★ 全文件最关键的一行：签名只对【前两段】做，签名自己放最后。
	//   顺序不能反——签名字段不能参与自己的签名，否则验签时无法自洽（循环依赖）。
	signing := h + "." + p
	return signing + "." + base64.RawURLEncoding.EncodeToString(hmacSHA256(signing, jwtSecret)), nil
}

// verifyToken：验签 + 过期检查。对应 golang-jwt 的 ParseWithClaims。
// 真实项目里"登录态校验中间件"的核心，其实就是这一个函数。
func verifyToken(token string) (map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		// 先做最便宜的格式检查，别浪费时间算哈希
		return nil, fmt.Errorf("格式错误：应该是 3 段，实际 %d 段", len(parts))
	}
	// ★ 验签的核心动作：用同一把密钥，对"第一段+第二段"重算一次签名，再和 token 的第三段比较
	expect := base64.RawURLEncoding.EncodeToString(hmacSHA256(parts[0]+"."+parts[1], jwtSecret))
	if !hmac.Equal([]byte(expect), []byte(parts[2])) {
		// 用 hmac.Equal 而不是 ==：它是恒定耗时比较，避免"时间侧信道攻击"
		// （普通 == 遇到第一个不同字节就返回，攻击者能靠响应耗时一位位猜出正确签名）
		return nil, fmt.Errorf("签名不匹配：payload 被篡改过")
	}
	// —— 到这里只证明了"没被改过"，还说明不了"没过期"，所以下面单独查 ——
	var payload map[string]any
	dec, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("payload 解码失败")
	}
	if err := json.Unmarshal(dec, &payload); err != nil {
		return nil, fmt.Errorf("payload 不是合法 JSON")
	}
	// exp 是 JWT 规范里的保留字段名（不是我们发明的），单位是 Unix 【秒】不是毫秒
	// 这里用 float64 断言，是因为 json.Unmarshal 进 map[string]any 时，数字一律解成 float64
	if exp, ok := payload["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
		return nil, fmt.Errorf("token 已过期")
	}
	return payload, nil
}

func jwtLogin(c *gin.Context) {
	token, err := signToken(map[string]any{
		"sub":  "zhangsan",                       // sub = subject，JWT 标准保留字段，表示"这个 token 是签给谁的"
		"role": "user",                           // 业务自定义字段。⚠️ 别往里塞密码、身份证这类敏感信息——payload 是明文的
		"exp":  time.Now().Add(time.Hour).Unix(), // 过期时间，Unix 秒
	})
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	// 传 token 的两种路子，各有所长（面试常问，要能说清取舍）：
	//   放 Cookie：浏览器自动带，前端零成本；但跨站请求也会自动带 → 必须自己防 CSRF
	//   放 Header：JS 手动带上，天然跨域、不受 SameSite 限制；代价是前端得自己存
	//              （存 sessionStorage 比 localStorage 好一些，但 XSS 之后两者都能被读，
	//               真正的解法是 CSP 和内容清理，不是换存储）
	c.SetCookie("jwt", token, int(time.Hour/time.Second), "/", "localhost", false, false)
	// ↑ 注意这里 httponly=false，纯粹是演示时方便你从前端把它打出来。生产必须是 true。
	c.JSON(200, gin.H{
		"msg":   "登录成功",
		"token": token,
		"提示":    "第一段第二段用 base64 编码而不是加密，所以 payload 内容一眼就能看懂 —— JWT 防篡改，不防偷看",
	})
}

func jwtMe(c *gin.Context) {
	// 业界惯例是 Authorization: Bearer <token>（RFC 6750）。真实项目只留这一条，别两个都收。
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ") // 剥掉固定前缀；没有前缀时 TrimPrefix 原样返回，安全
	if token == "" {
		token, _ = c.Cookie("jwt") // 没有 Header 就退回读 Cookie
	}
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录：没有 token", "hint": "先访问 /jwt/login"})
		return
	}
	payload, err := verifyToken(token)
	if err != nil {
		// 注意：生产环境不要区分告诉前端"签名错"还是"已过期"——那是给攻击者的提示。这里为了演示才返回原因
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "token 无效", "原因": err.Error()})
		return
	}
	// payload 里只有我们自己签进去的字段；要用户详情得去查库，别指望 token 里全有
	c.JSON(200, gin.H{"msg": "已登录", "payload": payload, "实例": instanceID})
}

// 考点演示：篡改 payload 后再验签，看它怎么被识破。
// 这个函数把"攻击者视角"完整走了一遍，是理解"JWT 防篡改"最直接的教材。
func jwtTamper(c *gin.Context) {
	good, _ := signToken(map[string]any{"sub": "zhangsan", "role": "user", "exp": time.Now().Add(time.Hour).Unix()})
	// 攻击者的三步：① base64 解出 payload（解得出，因为那是编码不是加密）
	//                ② 把 role 改成 admin 再编码回去
	//                ③ 签名那一段【原封不动抄回来】
	payload, _ := json.Marshal(map[string]any{"sub": "zhangsan", "role": "admin", "exp": time.Now().Add(time.Hour).Unix()})
	// 纯字符串拼接：好 token 第一段 + 伪造的第二段 + 好 token 第三段
	// 新手最容易卡在这："为什么不能这样拼接？" —— 因为验签时是对"第一段+第二段"重算签名，
	// 第二段一换，重算结果和抄来的第三段就对不上，直接被拒。
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
	// 所以整条链路的安全性只压在一件事上：jwtSecret 不能泄露。
	// 这也是为什么上面那个常量在生产必须从环境变量读、绝对不能进 git。
}