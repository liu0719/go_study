package main

import (
	"encoding/json"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	// 自己写的cors配置
	// r.Use(AddcorsHeader)

	// 用cors包达到同样的效果
	config := cors.DefaultConfig()
	// 添加允许源数组
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	//添加允许的请求方法
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH", "TRACE"}
	// 添加允许的请求头字段
	config.AllowHeaders = []string{"Content-Type", "Authorization", "school"}
	// 添加允许暴露的头部
	config.ExposeHeaders = []string{"school", "expose"}
	// 设置多长时间内不需要再预检
	config.MaxAge = time.Second * 10
	r.Use(cors.New(config))

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
	/*
		用cors包时不需要这个
		r.OPTIONS("/stu", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok!",
			})
		})
	*/

	// jsonp来处理跨域问题
	teacher := gin.H{
		"name":   "李老师",
		"age":    30,
		"gender": "女",
	}
	// c.String只能返回字符串，要想将json化的数据传回前端，要先将数据序列化为byte[],再套string转为字符串
	data, _ := json.Marshal(teacher)
	r.GET("teacher", func(c *gin.Context) {
		// 获取前端传来的函数名，返回时以字符串的格式用函数名的形式包裹数据返回给前端，前端会自动执行
		funcName := c.Query("func_name")
		// 写好字符串格式，前端会自动执行
		c.String(200, funcName+"("+string(data)+")")
	})
	r.Run(":80")
}
