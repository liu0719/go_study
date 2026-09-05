package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
)

// 异常处理

// 向外层传递错误
func div(a, b int) (c int, err error) {
	if b == 0 {
		// 创建一个错误,并返回
		err = errors.New("除数不能为零")
		return
	}
	return a / b, err
}
func serve() (err error) {
	res, err := div(3, 0)
	if err != nil {
		fmt.Println("运算结果：", res)
		return err
	}
	fmt.Println("服务结果为:", res)
	return
}

// 中断程序
func init() {
	_, err := os.ReadFile("XXX")
	if err != nil {
		// panic(err.Error())
		// log日志
		// log.Fatalln("出现错误")
		// 这种一般是用于初始化，一旦初始化出现错误，程序继续走下去也意义不大了，还不如中断掉
		// panic("出现错误")
	}
}

// 恢复程序
// 我们可以在一个函数里面，使用一个defer，可以实现对panic的捕获，以至于出现错误不至于让程序直接崩溃，这种一般也是框架层的异常处理所做的
func read() {
	// defer中的recover会捕获当前函数的异常，让程序能够走下去而不出错
	defer func() {
		// 捕获，也可以不捕获，交给上一层函数来捕获，直到main为止
		if err := recover(); err != nil {
			fmt.Println("read有错误,堆栈信息如下：")
			fmt.Println(string(debug.Stack()))
		}
	}()
	var list []int = []int{1, 2}
	fmt.Println(list[2]) //这里会产生一个越界panic
}
func main() {

	// 向外部传递，最外层接收在决策
	err := serve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("成功")
	}

	// 中断程序
	fmt.Println("我是中断程序")

	// 恢复程序
	read()
}
