package main

import (
	"encoding/json"
	"fmt"
	 "io"
	"net/http"
)

// 原生http包缺点
// 不区分请求，get 和post都可以访问
// 参数解析验证麻烦
// 请求处理比较原始
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		// 拿到请求体数据
		bytedata, _ := io.ReadAll(r.Body)
		fmt.Println(string(bytedata))
	}
	// 拿请求头
	fmt.Println(r.Header)

	// 返回json
	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}
	bytedata, _ := json.Marshal(Response{
		Code: 200,
		Msg:  "成功",
		Data: "这是数据",
	})
	w.Write(bytedata)
}

// go原生http包实现一个简单的web服务
func main() {

	http.HandleFunc("/index", IndexHandler)
	http.ListenAndServe(":80", nil)
}
