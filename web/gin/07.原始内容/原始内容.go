package main

import (
	"bytes"
	"fmt"
	"io"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		// 这个body本质上是从网络中拿的，阅后即焚，再读就读不到了
		byteData, _ := io.ReadAll(c.Request.Body)
		fmt.Println("===============================")

		fmt.Println(string(byteData))
		fmt.Println("================================")
		// 解决Body阅后即焚的问题
		// 只有将body的内容加载到到内存,再用bytes的NewReader读取内存中的bytedata
		// 再用io.NopCloser()将*reader包装成关闭空操作的同类型，再赋值给Body，就能实现反复读取
		c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
		fmt.Println(c.Request.Header)
		values,err:=url.ParseQuery("?")
		if err!=nil{
			fmt.Println(err)
		}
		for k, v := range values {
			fmt.Println(k,v)
		}
	})
	// 启动
	r.Run(":80")
}
