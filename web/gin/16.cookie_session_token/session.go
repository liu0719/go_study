package main

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// session
type Session struct {
	UserId   string
	LoginAt  time.Time
	ExpireAt time.Time
}
type MemorySessionStore struct {
	mu   sync.RWMutex
	data map[string]*Session
}

// 初始化
func NewMemorySessionStore() *MemorySessionStore {

	return &MemorySessionStore{data: make(map[string]*Session)}
}
func (m *MemorySessionStore) Get(sid string) (*Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	se, ok := m.data[sid]
	if !ok {
		return nil, false
	}
	// 如果现在的时间在指定的过期时间之后
	// 调用者时间是否在参数的时间之后
	if time.Now().After(se.ExpireAt) {
		m.mu.Lock()
		delete(m.data, sid)
		m.mu.Unlock()
		return nil, false
	}
	return se, true
}
func (m *MemorySessionStore) Set(sid string, se *Session) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[sid] = se
}

func (m *MemorySessionStore) Delete(sid string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, sid)
}
func (m *MemorySessionStore) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// 只统计非游客，且未过期的
	n := 0
	for _, se := range m.data {
		if se.UserId == "" || time.Now().After(se.ExpireAt) {
			continue
		}
		n++
	}
	return n
}

// 生成一个随机数ID,
func newId() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	// 返回base64编码转的string
	return base64.RawURLEncoding.EncodeToString(b)
}
func SessionMiddleWare(store *MemorySessionStore) func(*gin.Context) {
	return func(c *gin.Context) {
		// 将MemorySessionStore挂入Context
		c.Set("store", store)
		// 获取cookie
		sid, err := c.Cookie("sid")
		// 用cookie去store内获取Session
		_, ok := store.Get(sid)
		// 如果拿不到cookie或者和session核对失败
		if !ok || err != nil {
			// 新建一个sessionId
			sid := newId()
			// 存到store内
			// 中间件登录时给的只是游客模式
			store.Set(sid, &Session{
				UserId:   "",
				LoginAt:  time.Now(),
				ExpireAt: time.Now().Add(time.Minute * 30),
			})
			// 存储到cookie
			c.SetCookie("sid", sid, int(time.Minute*30/time.Second), "/session", "127.0.0.1", false, true)
		}
		// 拿到了就直接挂到Context上,
		c.Set("sid", sid)
		c.Next()

	}
}

// 登陆时创建Session
func SessionLogin(c *gin.Context) {
	store := c.MustGet("store").(*MemorySessionStore)
	sid := c.MustGet("sid").(string)
	store.Set(sid, &Session{
		UserId:   "张三",
		LoginAt:  time.Now(),
		ExpireAt: time.Now().Add(time.Minute * 30),
	})
	// 这里sessionId不变只是用户登陆了，更新一下cookie
	c.SetCookie("sid", sid, int(30*time.Minute/time.Second), "session", "127.0.0.1", false, true)
	c.String(200, "session设置成功")
}

// 判断当前状态
func SessionMe(c *gin.Context) {
	store := c.MustGet("store").(*MemorySessionStore)
	sid := c.MustGet("sid").(string)
	se, ok := store.Get(sid)

	// 调试
	// fmt.Printf("当前数量：%v,本地列表:\n", store.Count())
	// for _, v := range store.data {
	// 	fmt.Println(v.UserId)
	// }
	// 判断未登录
	if !ok || se.UserId == "" {
		c.String(200, "未登录")
		return
	}
	// 登陆就返回状态
	c.JSON(200, gin.H{
		"msg":    "已登录",
		"userID": se.UserId,
		"实例":     instanceID,
		"本机在线数":  store.Count(),
		// 转为指定格式时间
		"到期时间": se.ExpireAt.Format(time.RFC3339),
	})
}

// 登出时删除session，
func SessionLogout(c *gin.Context) {
	store := c.MustGet("store").(*MemorySessionStore)
	sid := c.MustGet("sid").(string)
	// 实现踢人功能
	store.Delete(sid)
	// 设置cookie时间，将其删掉
	c.SetCookie("sid", "", -1, "session", "127.0.0.1", false, true)
	c.JSON(200, "已退出登录")
}


func SessionGroup(r *gin.RouterGroup) {
	r.GET("login", SessionLogin)
	r.GET("me", SessionMe)
	r.GET("logout", SessionLogout)
}
