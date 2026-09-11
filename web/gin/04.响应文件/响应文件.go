package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// 在消息头设置内容类型，表示是文件流，唤起浏览器下载
		c.Header("Content-Type", "application/octet-stream")
		// 再设置文件名
		c.Header("Content-Disposition", "attachment; filename=photo.jpg")

		c.File("static/photo.jpg")
		c.JSON(200, gin.H{
			"msg": "文件正在下载",
		})
	})

	r.Run(":80")
}
