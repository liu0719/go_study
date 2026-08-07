package main

import (
	"errors"
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

// orm案例
type ClassModels struct {
	Id   int    `orm:"id"`
	Name string `orm:"-"`
}

// 模拟orm
func Find(obj any, query ...any) (err error, sql string) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	// 判断是否是结构体
	if v.Kind() != reflect.Struct {
		err = errors.New("不是结构体")
		return
	}
	// 如果有查询条件，就拼接where语句
	var where string
	// 判断参数个数和查询数能不能对的上
	if len(query) > 0 {
		//query[0]是查询条件,query[1:]是查询条件的值
		q := query[0]
		// 断言为字符串，如果不是字符串，就报错,qs方便后续使用
		qs, ok := q.(string)
		if !ok {
			err = errors.New("不是字符串,无法生成sql语句")
			return
		}
		// 统计?的个数,如果?的个数和参数个数不一致，就报错
		num := strings.Count(qs, "?")
		if num != len(query)-1 {
			err = errors.New("参数个数对不上")
			return
		}

		// for循环遍历query[1:]，把query[1:]的值替换成?
		for _, v := range query[1:] {
			switch s := v.(type) {
			case string:
				// replace,参数1:总字符串,参数2:要替换的字符串,参数3:要换成的字符串,0参数4:替换几个
				qs = strings.Replace(qs, "?", fmt.Sprintf("'%s'", s), 1)
			case int:
				qs = strings.Replace(qs, "?", fmt.Sprintf("'%d'", s), 1)
			default:
				err = errors.New("参数类型不支持")
				return
			}
		}
		// 拼接where语句
		where = "where " + qs
	}
	// 拼接sql语句
	// aim是要查询的字段
	aim := ""
	for i := 0; i < t.NumField(); i++ {
		// orm:"-" 表示不查询这个字段
		if t.Field(i).Tag.Get("orm") == "-" {
			continue
		}
		// 有orm标签，就拼接orm标签的值，转小写
		aim += strings.ToLower(t.Field(i).Name) + " "
	}
	// 表名是结构体名称的首字母小写+结构体名称的后5个字母小写
	table := strings.ToLower(t.Name()[:5]) + "_" + strings.ToLower(t.Name()[5:])
	// 拼接sql语句
	sql = fmt.Sprintf("select %s from %s %s", aim, table, where)
	return
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

	// orm例子
	err, sql := Find(ClassModels{}, "name =?", "三年一班")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql)
	// select name,id from class_models where name="三年一班"
	err, sql = Find(ClassModels{}, "id=? and name=?", 1, "三年一班")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql)
	// select name id from class_models where id =1 name="三年一班"
	err, sql = Find(ClassModels{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql)
	// select name id from class_models

}
