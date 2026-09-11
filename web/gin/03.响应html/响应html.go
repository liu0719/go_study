package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化
	r := gin.Default()

	// 对html文件进行预加载
	//
	r.LoadHTMLGlob("templates/*")
	//这个loadhtmlfiles要传入具体的文件路径
	r.LoadHTMLFiles("templates/index.html")

	r.GET("/", func(c *gin.Context) {
		// 响应html
		// 第二个参数只是单纯的文件名，不是路径
		// 第三个参数时要传给前端的数据,修改网站标题非常好用
		c.HTML(200, "index.html", map[string]any{
			"title": "我是标题",
		})
	})
	// 启动
	r.Run(":80")
}
