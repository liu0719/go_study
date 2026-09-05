package main

import "fmt"

func main() {
	// 打包命令，有些第三方包不能编译，只能在固定平台使用
	// -o 指定输出文件名
	// go build -o hello.exe main.go
	fmt.Println("hello world")
	// 再次注意啊，以后打包web项目的时候，配置文件和静态文件等这些非go程序，是要一起复制到目标服务器里面的

}
