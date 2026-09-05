package main

import (
	"fmt"
	"net/http"
	"os"
)

// w是响应对象，r是请求对象
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.UserAgent())
	data, err := os.ReadFile("19.netcoding\\http_server\\index.html")
	if err != nil {
		w.Write([]byte("文件不存在"))
		panic(err)
	}
	w.Write(data)
}
func Icon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("icon"))
}
func main() {

	http.HandleFunc("/", Index)
	// listenAndServe是http,ListenAndServeTLS是https
	// 第二个参数是handler,就是服务回调，做反向代理
	fmt.Println("监听127.0.0.1:80中")
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		panic(err)
	}

}
