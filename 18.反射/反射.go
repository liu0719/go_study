package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 反射
// 就比如数据库查询框架，不知道具体字段类型，就需要反射来获取字段类型，并判断变量是否是结构体，切片，map

// 获取类型
func GetType(obj any) {
	// 返回值不是真的类型，而是reflect.Type类型
	t := reflect.TypeOf(obj)
	// t.Kind() 返回值才是具体类型
	switch t.Kind() {
	case reflect.Struct:
		fmt.Println("是结构体")
	case reflect.Slice:
		fmt.Println("是切片")
	case reflect.Map:
		fmt.Println("是map")
	default:
		fmt.Println("是其他类型")
	}
}

// 获取值
func GetValue(obj any) {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("是整数", v.Int())
	case reflect.String:
		fmt.Println("是字符串", v.String())
	case reflect.Struct:
		fmt.Println("是结构体", v.Field(0))
	case reflect.Slice:
		fmt.Println("是切片", v.Index(0))
	case reflect.Map:
		fmt.Println("是map", v.MapKeys())
	default:
		fmt.Println("是其他类型", v.Interface())
	}
}

// 修改值
func SetValue(obj any, value any) {
	v1 := reflect.ValueOf(obj)
	v2 := reflect.ValueOf(value)
	// v1传过来是指针，所以需要v1.Elem()获取指针指向的值,相当于解引用
	if v1.Elem().Kind() != v2.Kind() {
		fmt.Println("类型不同，不能修改值")
		return
	}
	switch v1.Elem().Kind() {
	case reflect.Int:
		// 反射修改，加类型断言
		v1.Elem().SetInt(v2.Int())
	case reflect.String:
		v1.Elem().SetString(v2.String())
	case reflect.Struct:
		v1.Elem().Field(0).Set(v2.Field(0))
	case reflect.Slice:
		v1.Elem().Index(0).Set(v2.Index(0))
	case reflect.Map:
		v1.Elem().SetMapIndex(v2.MapKeys()[0], v2.MapIndex(v2.MapKeys()[0]))
	default:
		v1.Elem().Set(v2)
	}
}

// 通过反射操作结构体
func ParseStruct(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	// 根据传入的obj类型，判断是否是结构体
	if v.Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}
	for i := 0; i < v.NumField(); i++ {
		// 获取字段值
		field := t.Field(i)
		// 获取json标签，如果没有json标签，就使用字段名
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		// tf.Tag是结构体字段的标签，v.Field(i)是字段值
		fmt.Printf("结构体名字：%#v,类型：%#v,json：%#v,值：%#v\n", field.Name, field.Type.Name(), jsonTag, v.Field(i))
	}
}

// 修改结构体的值
func SetStructValue(obj any) {
	fmt.Println("-----修改值内容-----")
	// 获取结构体的值和类型,Elem()是解引用传过来是指针需要用，获取值
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	// 遍历结构体的字段
	for i := 0; i < v.NumField(); i++ {
		// v.Field(i)是结构体字段的值
		// t.Field(i)是结构体字段的类型

		// t.Field(i)是结构体字段的类型.Tag.Get("big")是获取big标签的值
		bigTag := t.Field(i).Tag.Get("big")
		if bigTag == "" {
			continue
		}
		// 修改字段值
		v.Field(i).SetString(strings.ToUpper(v.Field(i).String()))
	}
}

// 调用结构体的方法
type User struct {
	Name string
	Age  int
}

func (u User) Call(value string) {
	fmt.Println(u, value)
}
func Call(obj any) {
	// v里面都是方法，实例，动态的，用来执行的
	v := reflect.ValueOf(obj).Elem()
	// t的都是都是配置，静态的，用来看的
	t := reflect.TypeOf(obj).Elem()

	for i := 0; i < v.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		// if m.Name == "Call" {
		// 	continue
		// }
		method := v.Method(i)
		// Call方法需要reflect包里的Value类型的切片，初始化后，用ValueOf方法把要传的参数包装成ref包内的value类型
		method.Call([]reflect.Value{
			reflect.ValueOf("你好"),
		})
	}
}
func main() {
	// 获取类型
	GetType(123)
	GetType([]int{1, 2, 3})
	GetType(map[string]string{})
	// 简单的结构体定义
	GetType(struct{ name string }{name: "张三"})
	// 获取值
	GetValue(123)
	GetValue([]int{1, 2, 3})
	GetValue(map[string]string{})
	GetValue(struct {
		name string `json:"name"`
	}{name: "张三"})

	// 修改值
	var a int = 123
	fmt.Println("修改前a =", a)
	// 这里的a是值拷贝，所以需要传递指针
	SetValue(&a, 456)
	fmt.Println("修改后a =", a)

	// 通过反射操作结构体
	ParseStruct(struct {
		name  string `json:"name"`
		age   int
		IsMan bool
	}{name: "张三", age: 25, IsMan: true})

	s1 := struct {
		Name1 string `big:"-"`
		Name2 string
	}{Name1: "name1", Name2: "name2"}
	SetStructValue(&s1)
	fmt.Println(s1)
	user1 := User{
		Name: "张三",
		Age:  25,
	}
	Call(&user1)
}
