package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个测试 Gin 应用
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		// 第一次读取请求体
		byteData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, "读取请求体失败")
			return
		}
		log.Printf("第一次读取到的数据: %s", string(byteData))

		// 重新包装请求体，使其可以再次读取
		c.Request.Body = io.NopCloser(bytes.NewReader(byteData))

		// 第二次读取请求体
		byteData2, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, "第二次读取请求体失败")
			return
		}
		log.Printf("第二次读取到的数据: %s", string(byteData2))

		// 第三次读取请求体（再次重新包装）
		c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
		byteData3, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, "第三次读取请求体失败")
			return
		}
		log.Printf("第三次读取到的数据: %s", string(byteData3))

		c.String(http.StatusOK, "成功读取请求体3次")
	})

	// 创建测试请求
	reqBody := "这是测试请求体数据"
	req, _ := http.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "text/plain")

	// 发送请求
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	fmt.Printf("响应状态码: %d\n", w.Code)
	fmt.Printf("响应内容: %s\n", w.Body.String())
}