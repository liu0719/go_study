package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 静态文件路径代替，第一个参数是url路径(别名)，第二个是文件路径
	r.Static("st", "static")
	// 这个只能用于单个文件
	r.StaticFile("hello", "static/hello.txt")

	r.Run(":80")

}
