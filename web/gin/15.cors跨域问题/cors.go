package main

import "github.com/gin-gonic/gin"

// 这种中间件已经有人封装了，就是cors包
func AddcorsHeader(c *gin.Context) {

	// http响应头必须要在发送响应内容前设置。
	// 也就是说要在处理响应内容前的中间件设置好,不能在c.Next()后,此时响应内容已经发送走
	// 设置响应头的value字段必须严格相等,不然还是禁止跨域
	c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:5500") //添加允许跨境的源， "*"为全部允许

	// 应对复杂请求,复杂请求浏览器都会有预检请求
	// 设置响应头
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH,TRACE") // 添加允许跨境的方法
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization,school")           //添加允许跨境的自定义请求头
	c.Header("Access-Control-Max-Age", "10000")                                              //设置多长时间内不需要再预检
}
func main() {
	r := gin.Default()
	r.Use(AddcorsHeader)
	// 简单请求
	r.GET("/stu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":   "张三",
			"age":    18,
			"gender": "男",
			"score":  88.88,
		})
	})
	// 复杂请求
	r.DELETE("/stu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":   "张三",
			"age":    18,
			"gender": "男",
			"score":  88.88,
		})
	})
	// 应对跨域复杂请求的预检options请求方式，必须有成功返回值
	r.OPTIONS("/stu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok!",
		})
	})
	r.Run(":80")
}
