package main

// 数组切片map

import (
	"fmt"
	"slices"
	"sort"
)

func main() {

	// 数组定义
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	var arr2 = [3]float32{1, 2, 3}
	fmt.Println(arr2)
	var arr3 = [...]string{"hello", "world", "golang"}
	fmt.Println(arr3)

	// 修改
	arr1[2] = 99
	fmt.Println(arr1)

	// 数组长度是死的，不能改，切片可以
	var list1 []string
	list1 = append(list1, "你好")
	list1 = append(list1, "我叫张三")
	fmt.Println(list1)
	list1[0] = "我也不知道"
	fmt.Println(list1)

	// make函数
	// 除了基本数据类型，其他类型不赋值,默认值就是nil
	// 定义一个字符串切片,默认是空
	var list2 []string
	fmt.Println(list2 == nil) // true

	//格式： make([]type, length, capacity)
	list3 := make([]string, 5, 10)
	list3[0] = "h"
	list3[1] = "e"
	list3[2] = "l"
	list3[3] = "l"
	list3[4] = "o"
	fmt.Println(list3, len(list3), cap(list3))

	fmt.Println(list3 == nil) //flase
	// 切的是左闭右开区间
	slices1 := list3[1:]
	fmt.Println(slices1)

	// 切片排序
	var list4 = []int{4, 99, 23, 12, 45, 76, 32}
	//引用排序库
	sort.Ints(list4)
	fmt.Println("升序", list4)
	// sort.Sort(sort.Reverse(sort.IntSlice(list4)))
	slices.Reverse(list4)
	fmt.Println("降序", list4)

	// map
	// key-value类型，key必须为基本数据类型
	// 声明,map一定要赋值，只声明是无效的，会报错
	var m1 map[string]string
	// 初始化1
	m1 = make(map[string]string)
	// 初始化2
	m1 = map[string]string{}
	// 赋值
	m1["name"] = "张三"
	fmt.Println(m1)
	// 取值
	fmt.Println(m1["name"])
	// 删除
	delete(m1, "name")
	fmt.Println(m1)

	m2 := map[string]string{
		"age":     "你好",
		"gender":  "男",
		"name":    "李四",
		"address": "河北",
	}
	res := m2["age1"] //取一个不存在的
	fmt.Printf("%#v\n", res)
	res, ok := m2["age1"]
	fmt.Printf("res:%#v,ok:%#v\n", res, ok) //res为空，ok为false
	// 取存在的
	res, ok = m2["age"]
	fmt.Printf("res:%#v,ok:%#v\n", res, ok) //res为值，ok为true
	// 删除
	fmt.Println("删除前:", m2)
	delete(m2, "address")
	fmt.Println("删除后", m2)
	m2["work"] = "程序员"
	fmt.Println(m2)
}
