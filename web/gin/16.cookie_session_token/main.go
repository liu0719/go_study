package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.SigningMethod
}

func main() {
	r := gin.Default()

	// cookie组,不用中间件
	cookieG := r.Group("cookie")
	CookieGroup(cookieG)

	// session
	// 初始化一个store
	store := NewMemorySessionStore()
	sessionMW := SessionMiddleWare(store)
	r.Use(sessionMW)
	sessionG := r.Group("session")

	SessionGroup(sessionG)

	// token

	r.Run(":80")
}
