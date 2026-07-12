package main

// 输入输出

import (
	"fmt"
)

func main() {

	// 输出

	fmt.Println("hello", "你好")
	// %s字符串，%d整数，%.3浮点数，小数点后保留三位
	fmt.Printf("我叫%s,今年刚满%d岁,成绩%.3f分", "张三\n", 18, 20.34)
	// %T打印数据类型
	fmt.Printf("%T %T\n", "你好", 2.4)
	// %v 万能打印符
	fmt.Printf("%v\n", "你好")
	// %#v,带#号可以把引号打印出来，鉴别空字符串
	fmt.Printf("%#v\n", "")

	// 输入
	fmt.Print("请输入文本\n")
	var name string = ""
	n, err := fmt.Scanf("%s", &name) //n为布尔，err报错内容
	fmt.Printf("你输入的是{%v}\n", name)
	fmt.Println(n, err)
}
