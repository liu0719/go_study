package main

import "fmt"

// 基本数据类型

func main() {

	// 整形，
	// 后面带的数如8，16，32，64，就是二进制存储位数
	// uint也叫byte,int取决于设备的硬件
	// 带u为无符号数，仅正数
	var n1 uint8 = 2
	var n2 uint16 = 2
	var n3 uint32 = 2
	var n4 uint64 = 2
	var n5 uint = 2
	// 有符号数，有正负
	var n6 int8 = 2
	var n7 int16 = 2
	var n8 int32 = 2
	var n9 int64 = 2
	var n10 int = 2
	fmt.Println(n1, n2, n3, n4, n5, n6, n7, n8, n9, n10)

	// 浮点数,32,64,64位的更长
	var f1 float32 = 3.253
	var f2 float64 = 3.547392463
	fmt.Println(f1, f2)

	// 字符型
	// 字符的本质还是一个数
	// byte（uint8)是单字节字符，rune是多字节字符，32位
	// uint8就是ASCll编码
	var c1 = 'a'
	var c2 = 97
	fmt.Println(c1) // 直接打印都是数字
	fmt.Println(c2)

	fmt.Printf("%d %c\n", c1, c2) // 以字符的格式打印
	//rune可存储中文
	var r1 rune = '中'
	fmt.Printf("%d\n", r1)

	// 字符串
	// 字符串是双引号，字符型是单引号
	var s1 string = "你好啊"
	fmt.Println(s1)

	//转义字符
	fmt.Println("枫枫\t知道")              // 制表符
	fmt.Println("枫枫\n知道")              // 回车
	fmt.Println("\"枫枫\"知道")            // 双引号
	fmt.Println("枫枫\r知道")              // 回到行首
	fmt.Println("C:\\pprof\\main.exe") // 反斜杠

	//多行字符串

	var s = `今天
		    天气真好`
	fmt.Println(s)

	// 布尔值
	var flag = true
	flag = false
	fmt.Println(flag)

	//只声明不赋值，自动默认值，0，false,"",
	var a1 int
	var a2 float32
	var a3 string
	var a4 bool

	fmt.Printf("%#v\n", a1)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)

}
