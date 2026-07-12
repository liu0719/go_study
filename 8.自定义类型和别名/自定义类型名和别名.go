package main

import "fmt"

// 自定义类型名和别名

// 类型别名和原类型区别
// 原类型不能绑定方法
// 打印类型还是原始类型

type Code int
type MyCode int

const (
	SUCCESSCODE Code = 0
	SERVICEERR  Code = 101 //验证失败响应码
	NETWORKERR  Code = 201 //错误响应码
	
)
// 根据code的不同类型返回不同结果
func (c Code) getMsg() string {
	switch c {
	case SUCCESSCODE:
		return "成功"

	case SERVICEERR:
		return "服务器错误"
	case NETWORKERR:
		return "网络错误"
	}
	return "未知错误"
}
// 直接返回两个值
func (c Code)ok()(Code,string){
	return c,c.getMsg()
}
func webserve(name string) (code Code, msg string) {
	if name == "1" {
		return SERVICEERR.ok()
	}
	if name == "2" {
		return NETWORKERR.ok()
	}
	return SUCCESSCODE.ok()
}

func main() {
	const age  =1
	var c Code=1
	fmt.Println(c==age)
}
