package main

import (
	"encoding/json"
	"fmt"
)

// 泛型

// 泛型函数
// 可以自定义类型,表示多种类型的接口，写的时候方便
// 类型集合，表示允许的类型集合
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// 在函数名后用[T int|float32]表示T为多种类型
func add[T Number | float32 | uint](a, b T) T {
	return a + b
}
func Myadd[T Number | float32 | uint, K string | int](a T, b K) K {
	return b
}

// 泛型结构体
type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	// 要用T指代，就能实现套着好几层连点
	Data T `json:"data"`
}

func main() {
	type User struct {
		Name string `json:"name"`
	}
	type Userinfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user := Response[User]{
		Code: 200,
		Msg:  "成功",
		Data: User{
			Name: "张三",
		},
	}
	// 表明内部嵌套的结构，就可以连点
	userinfo := Response[Userinfo]{
		Code: 200,
		Msg:  "成功",
		Data: Userinfo{
			Name: "张三",
			Age:  32,
		},
	}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
	byteData, _ = json.Marshal(userinfo)
	fmt.Println(string(byteData))

	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":200,"msg":"成功","data":{"name":"张三"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)
	var userinfoResponse Response[Userinfo]
	json.Unmarshal([]byte(`{"code":200,"msg":"成功","data":{"name":"张三","age":32}}`), &userinfoResponse)
	fmt.Println(userinfoResponse.Data.Name, userinfoResponse.Data.Age)

	// 泛型切片
	type Myslice[T int | string] []T
	// 切片声明的时候要指定类型
	var myslice = Myslice[int]{1, 2, 3, 4}
	fmt.Println(myslice)

	// 泛型map
	type MyMap[K string, V any] = map[K]V

	mymap := MyMap[string, any]{
		"name":  "张三",
		"age":   10,
		"score": 21.345,
	}
	fmt.Println(mymap)
}
