package main

import (
	"encoding/json"
	"fmt"
)

// 结构体

// type 结构体名 struct{}
type Person struct {
	// 结构体tag
	// json只会序列化能导出的属性，所以属性想序列化要大写
	// 首字母大写 类型 `json:"json化后的名字"`  不想json化直接用"-"
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Score   float32 `json:"score,omitempty"` //"score,omitempty"忽略空值
	Address string  `json:"address"`
	Gender  string  `json:"gender"`
	Work    string  `json:"-"` //"-"表示忽略该段
}

// 结构体上的方法
func (p *Person) PrintInfo() {
	fmt.Printf("姓名：%v，性别：%v,年龄%v，地址：%v，成绩：%v，工作：%v。\n", p.Name, p.Gender, p.Age, p.Address, p.Score, p.Work)
}

// 结构体指针
func (p *Person) setName(name string) {
	p.Name = name
}
func (p *Person) getName() string {
	return p.Name
}

type Student struct {
	person Person //直接当属性，来实现继承
	stu_id int
}

func (s *Student) getName() string {
	// 有继承关系
	return s.person.Name
}

func main() {
	p := Person{
		Name:    "张三",
		Age:     20,
		Address: "河北",
		Gender:  "男",
	}
	p.PrintInfo()
	p.setName("王五")
	p.PrintInfo()
	byteData, _ := json.Marshal(p)
	fmt.Println(string(byteData))
	fmt.Println()
}
