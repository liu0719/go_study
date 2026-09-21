package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// api路径分组
func UserGroup(r *gin.RouterGroup) {
	r.GET("users", UserView)
	r.POST("users", UserView)
	r.PUT("users", UserView)
	r.PATCH("users", UserView)
	r.DELETE("users", UserView)
	r.HEAD("users", UserView)
}

// 同名api路径分组
func LoginGroup(r *gin.RouterGroup) {
	r.GET("login", LoginView)
}

// /api/users路径处理函数
func UserView(c *gin.Context) {
	path := c.Request.URL
	fmt.Println(path, c.Request.Method)
	c.JSON(200, gin.H{
		"msg": "成功",
	})
}

// /api/login路径处理函数
func LoginView(c *gin.Context) {
	path := c.Request.URL
	fmt.Println(path, c.Request.Method)
	c.JSON(200, gin.H{
		"msg": "LOGIN页面",
	})
}

func main() {
	r := gin.Default()
	// 1.路由
	/*
		r.GET() 下载资源
		r.POST() 提交资源
		r.PUT() 全面更新资源
		r.PATCH() 部分更新
		r.DELETE() 删除
		r.Any()  支持所有请求方式
	*/

	// 2.路径
	apiGroup := r.Group("api")
	// 分组方便管理,可以统一配置中间件,先走中间件再走各种路由方法
	apiGroup.Use()
	UserGroup(apiGroup)
	// 可以创建同名的路由,下面是一个不用中间件的路由
	noMiddleWarGroup := r.Group("api")
	LoginGroup(noMiddleWarGroup)

	r.Run(":80")
}
