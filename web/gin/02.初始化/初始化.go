package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liu0719/go_study/web/gin/res"
)

func Index(c *gin.Context) {
	fmt.Println(c.Request)

	// 3.响应json,gin.H是一个map[string]any类型
	// 响应码一般都是200，
	c.JSON(200, gin.H{
		// 这个code是业务码，判断业务层面是否正确
		"Code": 200,
		"Msg":  "后端数据",
		"Data": gin.H{},
	})
}

func main() {
	// 0.隐藏终端的gin日志
	gin.SetMode(gin.ReleaseMode)
	// 1.初始化，会返回一个gind的引擎，engine
	r := gin.Default()

	// 2.挂载路由
	r.GET("/", Index)

	r.GET("/login", func(c *gin.Context) {
		res.OkWithMsg(c, "登陆成功")
	})

	r.GET("/users", func(c *gin.Context) {
		res.OkwithData(c, map[string]any{
			"name": "张三",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		res.FailWithMsg(c, "验证失败")
	})
	r.Run(":80")
}
