package main

// 变量定义

import (
	"fmt"
	"go_study/version"
)

func hello() {
	fmt.Println()
	fmt.Println("你好")
}

// 常量声明
const (
	//变量属性方法要想实现跨包访问，首字母必须大写
	Pai = 3.14

	//Version="2.1.0"

)

// 全局变量不能用短声明符号，用var
// 全局变量不用不会爆红，函数内的不用会爆红
// go声明都是名称在前，类型在后

var age = 10

func main() {
	// 声明
	var name string
	// 赋值,类型可省略
	var name1 = "world"
	// 短声明符号，直接赋值
	name2 := "hello go"
	name = "world"
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println("hello go,", name)
	hello()

	// 多变量声明
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// 括号声明
	var (
		a1 = "1"
		b1 = "2"
		c1 = "3"
	)
	fmt.Println(a1, b1, c1)
	fmt.Println(version.Version)

}
