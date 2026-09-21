package main

import (

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 用环境变量区分实例，方便起两个进程体验"多实例丢登录态"
var (
	port       string
	instanceID string
)

func SetCookieDemo(c *gin.Context) {
	// 参数解析
	// 1.2.key-value：cookie的键值对
	// 3.有效时长
	// 4.那些路径下携带cookie
	// 5.域名设置。.exmple.com。二级子域名可共享
	// 6.是否只走https,加s安全
	// 7.是否只允许http操作，不允许前端js控制
	c.SetCookie("uid", "u_321", int(time.Minute*30/time.Second), "/cookie", "127.0.0.1", false, true)
	c.String(http.StatusOK, "cookie设置成功")
}
func GetCookieDemo(c *gin.Context) {
	cookie, err := c.Cookie("uid")
	if err != nil {
		c.String(200, err.Error())
	}
	c.String(200, fmt.Sprintf("获取cookie成功,cookie值为：%v", cookie))
}
// 路由组
func CookieGroup(r *gin.RouterGroup) {
	r.GET("/set", SetCookieDemo)
	r.GET("/get", GetCookieDemo)
}
