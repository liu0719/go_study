package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 查询参数的绑定,能直接绑定到指定结构体
	r.GET("/", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		// 也会完成数据的校验
		err := c.ShouldBindQuery(&user)
		fmt.Println(user, err)
	})
	// 动态参数的绑定
	r.GET("user/:id/:name/*action", func(c *gin.Context) {
		type User struct {
			Name   string `uri:"name"`
			Id     int    `uri:"id"`
			Action string `uri:"action"`
		}
		var user User
		err := c.ShouldBindUri(&user)
		fmt.Println(user, err)
	})
	// 表单参数
	/*
		r.POST("/user", func(c *gin.Context) {
			type User struct {
				Name string `form:"name"`
				Age  int    `form:"age"`
			}
			var user User
			// 也会完成数据的校验
			err := c.ShouldBind(&user)
			fmt.Println(user, err)
		})
	*/
	r.POST("json", func(c *gin.Context) {
		// 一定要大写，不然获取不到
		type User struct {
			Name    string `json:"name"`
			Age     int    `json:"age"`
			address string `json:"address"`
		}
		var user User
		// 也会完成数据的校验
		err := c.ShouldBindJSON(&user)
		fmt.Println("json数据为：", user, err)
	})
	r.POST("header", func(c *gin.Context) {
		// 一定要大写，不然获取不到
		type Header struct {
			UserAgent   string `header:"User-Agent"`
			ContentType string `header:"Content-Type"`
			Name        string `header:"name"`
		}
		var header Header
		// 也会完成数据的校验
		err := c.ShouldBindHeader(&header)
		fmt.Println("json数据为：", header, err)
	})
	r.Run(":80")
}
