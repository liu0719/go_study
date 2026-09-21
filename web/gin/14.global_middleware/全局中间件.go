package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

func ApiRouterGroup(r *gin.RouterGroup) {
	r.GET("users", func(c *gin.Context) {
		fmt.Println("Home")
		// 获取中间件的参数
		value, exist := c.Get("GM1")
		fmt.Println(value, exist)
		// 对获取到的参数进行类型断言
		_user, ok := c.Get("user")
		if ok {
			user, ok := _user.(User)
			if ok {
				fmt.Println(user.Name)
			}
		}
		c.String(200, "Home")
	})
	r.POST("users", func(c *gin.Context) {})
	r.PUT("users", func(c *gin.Context) {})
	r.DELETE("users", func(c *gin.Context) {})
}
func GM1(c *gin.Context) {
	fmt.Println("GM1请求")
	//再在这里设置的参数，在这个请求往后的生命周期都是可以拿到的，GM2,Home,GM2响应，GM1响应都能拿到
	c.Set("GM1", "我是GM1的数据")
	// 可以传任何参数，因为支持类型断言

	c.Set("user", User{Name: "不定积分"})
	c.Next()
	fmt.Println("GM1响应")
	value, exist := c.Get("GM1")
	fmt.Println(value, exist)
	c.String(200, "GM1响应")
}
func GM2(c *gin.Context) {
	fmt.Println("GM2请求")
	value, exist := c.Get("GM1")
	fmt.Println(value, exist)
	c.Set("GM2", "我是GM2的数据")
	// c.Abort()

	c.Next()
	fmt.Println("GM2响应")
	value, exist = c.Get("GM1")
	fmt.Println(value, exist)
	c.String(200, "GM2响应")
}

func AuthMiddleWare(c *gin.Context) {
	// 在这里可以拿header的token来确定权限，或者记录日志
	// 进行拦截或者放行
}

func main() {
	// Default方法其实就是调用New方法，然后使用了两个中间件logger(),recover(),分别负责日志，捕获panic.
	// 如果用New()就没有这些东西，需要自己写
	r := gin.Default()

	// 用Group方法创建新的路由分组
	g := r.Group("api")
	g.Use(GM1, GM2)
	// 再调用路由组函数处理新建的路由组。
	ApiRouterGroup(g)
	r.Run(":80")

}
