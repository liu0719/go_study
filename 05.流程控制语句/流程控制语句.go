package main

import "fmt"

// 流程控制语句

func main() {
	// if-else  , &&并 ,  ||或
	var age int
	fmt.Println("请输入你的年龄:")
	n, err := fmt.Scanf("%v", &age)

	if age > 0 && age < 18 {
		fmt.Println("未成年")
	} else if age >= 18 && age < 69 {
		fmt.Println("成年时期")
	} else if age >= 60 && age < 120 {
		fmt.Println("老年时期")
	} else if age < 0 || age >= 120 {
		fmt.Println("年龄不在规定范围内")
	} else if age == 0 || err != nil {
		fmt.Println("您的输入不合法")
	}
	fmt.Println(age)
	fmt.Println(n, err)

	fmt.Println("------------switch-------------")
	switch {
	case age == 0 && err != nil:
		fmt.Println("您的输入不合法")
	case age >= 0 && age < 18:
		fmt.Println("未成年")
	case age >= 18 && age < 60:
		fmt.Println("成年时期")
	case age >= 60 && age < 120:
		fmt.Println("老年时期")
	case age >= 120:
		fmt.Println("年龄不在规定范围内")
	}

	// 枚举 default
	var week int = 1
	fmt.Println("请输入1到7的数字")
	fmt.Scanf("%v", &week)
	switch week {
	case 1:
		fmt.Println("星期一")
		// 强制执行下一个case语句
		fallthrough
	case 2:
		fmt.Println("星期二")
		fallthrough
	case 3:
		fmt.Println("星期三")
		fallthrough
	case 4:
		fmt.Println("星期四")
		fallthrough
	case 5:
		fmt.Println("星期五")
		fallthrough
	case 6:
		fmt.Println("星期六")
		fallthrough
	case 7:
		fmt.Println("星期天")
		fallthrough
	default:
		fmt.Println("输入不符合规定")
	}

	// for
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
	s1 := make([]int, 0, 0)
	s1 = append(s1, 2)
	s1 = append(s1, 4)
	s1 = append(s1, 9)
	s1 = append(s1, 3)
	m1 := make(map[string]string)
	m1["address"] = "河北"
	m1["name"] = "张三"
	m1["gender"] = "女"
	m1["age"] = "13"
	//下标和内容
	for k, v := range s1 {
		fmt.Printf("k=[%#v],v=[%#v]\n", k, v)
	}
	// key和value
	for k, v := range m1 {
		fmt.Printf("k=[%#v],v=[%#v]\n", k, v)

	}

}
