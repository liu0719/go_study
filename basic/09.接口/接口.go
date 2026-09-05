package main

import "fmt"

// 接口

// 唱歌接口
// 实现接口：一个类型实现了接口的所有方法，即实现了该接口
type SingInterface interface {
	Sing()
	GetName() string //返回值也要写上，写全
}

// 鸡类型
type Chicken struct {
	Name string
}

// 鸡的两个方法，用来实现接口
func (c Chicken) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Chicken) GetName() string {
	return c.Name
}

// 鸭类型
type Duck struct {
	Name string
}

// 鸭的方法
func (d Duck) Sing() {
	fmt.Println(d.Name, "在唱歌")
}
func (d Duck) GetName() string {
	return d.Name
}

func sing(c SingInterface) {
	// 一种类型断言，oks是bool值
	// ch, ok := c.(Chicken)
	// fmt.Println(ch, ok)

	// c.Sing()

	// 全类型断言
	switch ser := c.(type) {
	case Chicken:
		fmt.Println("我是鸡：", ser)
	case Duck:
		fmt.Println("我是鸭：", ser)

	}

}

// 包含参数，方法名，返回值，但未具体实现的集合
// 接口本身不能绑定方法，接口是值类型，保存的是值+原始类型
type Animal interface {
	sing()
	jump()
	rap()
}

//任何类型都实现了空接口
type Empty interface{}

// 空接口便捷写法，传入任何类型都可以
// any就是interface{}，可以传入任何类型
func Print(Empty interface{}) {
	fmt.Println(Empty)
}
func main() {
	ch := Chicken{
		Name: "鸡哥",
	}
	du := Duck{
		Name: "鸭子",
	}
	sing(ch)
	sing(du)
}
