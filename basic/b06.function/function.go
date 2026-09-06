package main

import (
	"fmt"
	"time"
)

// 函数

func say_hello() {
	fmt.Println("hello")
}

// 函数有限个参，无返回值
func printlimit(a int, b int) {
	fmt.Println(a, b)
}

// 函数传多参，无返回值，
func print(nums ...[]int) {
	fmt.Println(nums)
}

// 函数传参 有限个返回值
func returnOne(a int) int {
	return a + 1
}

// 函数传参，多返回值
// 返回值可以提前定义，返回时只写return就行
func returnMore(a int) (val int, ok error) {
	return
}

// 闭包
// 当返回值的函数内用到了原函数的参数，就是闭包
// awaitAdd函数返回值是一个函数，返回的函数用来处理加法，在返回一个和就可以了
func awaitAdd(t int) func(...int) int {
	time.Sleep(time.Second * time.Duration(t))
	return func(nums ...int) (sum int) {
		for _, v := range nums {
			sum += v
		}
		return sum
	}
}

// 值传递
func copyName(name string) {
	fmt.Println(&name)
}

// 引用传递
// *string表示解引用，也可以理解为带*是内存地址  ，&为取地址符
func copyNameAddress(name *string) {
	fmt.Println(name)
}

// init函数
// 文件执行顺序   import包引用，const常量，var变量  init初始化，main函数
// init函数不能被调用，自动在main前用,init不能当参数被传入,init不能有传参或返回值
func init() {
	fmt.Println("初始化1")
}

// defer函数，在return前自动执行，如果有多个defer函数，先执行离return近的，再执行远的
// defer中的变量，在defer声明时就已经决定了，在defer后声明的变量，defer是拿不到的
func func_defer() {
	var age = 10
	defer fmt.Println("d1")
	fmt.Println("hello")
	defer fmt.Println("d2")
	defer func() {
		// 拿不到name
		fmt.Println(age)
		// 后面加()是为了调用该函数，如果不加只是声明了函数，没有调用
	}()
	var name string = "你好"
	fmt.Println(name)
}
func main() {
	//函数调用
	say_hello()
	//多参数，多返回值
	printlimit(1, 2)
	res := returnOne(1)
	fmt.Println(res)
	n, ok := returnMore(88)
	fmt.Println(n, ok)

	// 匿名函数
	// go不能在函数里再次创建函数，通过匿名函数解决
	var setName = func(name *string) {
		// 解引用，name传过来是个地址，加个*从*string变为string
		*name = "李四"
		fmt.Println(name)
	}
	var getName = func() string {
		return "张三"
	}
	name1 := "李四"
	setName(&name1)
	fmt.Println(getName())

	// 高阶函数
	// 就是把函数当作参数传到另一个参数内
	fmt.Println("--------------------高阶函数------------------------")

	fmt.Println("请输入你的操作：")
	fmt.Println(`1.登录
	2.注册
	3.游客登陆
	4.看戏`)
	var flag int
	fmt.Scan(&flag)
	var funcMap = map[int]func(){
		1: func() {
			fmt.Println("输入用户名和密码")
		},
		2: func() {
			fmt.Println("开始注册")
		},
		3: func() {
			fmt.Println("游客登陆功能有限")
		},
		4: func() {
			fmt.Println("不许看")
		},
	}
	funcMap[flag]()
	// 闭包掉用，awaitAdd(2)返回值是函数，也就是awaitAdd(2)是那个返回值函数名，而后面的（1，2，3）是返回值函数调用的参数
	fmt.Println(awaitAdd(2)(1, 2, 3))

	// 值传递
	name := "张三"
	fmt.Println(&name)
	copyName(name)

	// 引用传递
	fmt.Println(&name)
	copyNameAddress(&name)

	// defer函数
	func_defer()
}
