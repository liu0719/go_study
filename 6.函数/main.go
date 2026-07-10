package main

import "fmt"

// 函数
func say_hello() {
	fmt.Println("hello")
}

// 函数有限个参，无返回值
func printlimit(a int, b int) {
	fmt.Println(a, b)
}

// 函数传多参，无返回值
func print(nums ...[]int) {
	fmt.Println(nums)
}

// 函数传参 有限个返回值
func returnOne(a int) int {
	return a + 1
}

// 函数传参，多返回值
func returnMore(a int) (int, error) {
	return a + 1, nil
}

// han'x'chu
func main() {
	//函数调用
	say_hello()
	//多参数，多返回值
	printlimit(1, 2)
	res := returnOne(1)
	fmt.Println(res)
	n, ok := returnMore(88)
	fmt.Println(n, ok)
	

}
