package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

// trace的编译过程
// 1.创建文件
// 2.启动
// 3.停止
func main() {
	runtime.GOMAXPROCS(1) // 设置最大CPU数量为1
	// 创建trace文件
	f, err := os.Create("GMPmodel/trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 要调度的业务逻辑
	fmt.Println("Hello, World!")

	// 停止trace
	trace.Stop()
}
