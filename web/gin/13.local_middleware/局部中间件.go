package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("home请求")
	c.String(200, "home响应")
}
func M1(c *gin.Context) {
	fmt.Println("m1请求")
	// 拦截住，不会再往下走了，走完M1请求走M1响应
	// c.Abort()

	// Next方法会把函数分为两部分，强制进入下一个处理函数，就类似函数递归调用
	// 分为请求部分和响应部分，走完Home之后再走相应部分
	c.Next()
	c.String(200, "M1响应")
}
func M2(c *gin.Context) {
	// 请求部分在处理函数完成前，发送给服务器时处理
	fmt.Println("m2请求")

	// M1请求>M2请求>M2abort拦截，阻止往下走>M2响应>M1响应
	c.Abort()

	c.Next()
	// 响应部分在处理函数完成后，返回给客户端时再处理
	c.String(200, "M2响应")
}
func main() {
	r := gin.Default()
	// M1请求>Home>M1响应
	r.GET("m1", M1, Home)
	// M1请求>M2请求>Home>M2响应>M1响应，和主机之间7层王略结构通信类似
	r.GET("m1m2", M1, M2, Home)

	r.Run(":80")
}
