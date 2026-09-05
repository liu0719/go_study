package main

import "testing"

// 单元测试
// 1. 测试用例文件名必须以_test.go结尾
// 2. 测试用例函数必须以Test开头，一般来说就是Test+被测试的函数名

// 终端运行：go test  可以运行某个包下的所有测试用例
// -v  可以打印出详细的测试用例信息
// -run 参数可以指定测试某个函数

func TestAdd(t *testing.T) {
	// 单元测试框架提供的方法
	/*
			方 法	        备 注	                     测试结果
		    Log	     打印日志，同时结束测试	               PASS
		    Logf	格式化打印日志，同时结束测试           PASS
			Error	打印错误日志，同时结束测试	            FAIL
			Errorf	格式化打印错误日志，同时结束测试	     FAIL
		fatal会终止测试，后续的测试用例不会执行，在子测试中也会终止子测试的执行
			Fatal	打印致命日志，同时结束测试	            FAIL
		    Fatalf	格式化打印致命日志，同时结束测试	     FAIL
	*/
	// 测试用例
	if Add(1, 2) != 3 {
		t.Error("TestAdd() 测试失败")
	}
	if Add(4, -1) != 3 {
		t.Errorf("期望值：%d, 实际值：%d", 3, Add(4, -1))
	}
	t.Logf("期望值：%d, 实际值：%d", 3, Add(1, 2))
	t.Run("ADD1", func(t *testing.T) {
		if Add(1, 2) != 3 {
			t.Error("ADD1 测试失败")
		}
	})
	t.Run("ADD2", func(t *testing.T) {
		if Add(3, -2) != 1 {
			t.Error("ADD2 测试失败")
		}
	})
}
func TestSub(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Error("TestSub() 测试失败")
		t.Errorf("期望值：%d, 实际值：%d", -1, Sub(1, 2))
		// t.Fatal("TestSub() 测试失败，程序退出")
		t.Fatalf("期望值：%d, 实际值：%d", -1, Sub(1, 2))
		// t.FailNow() // 直接结束测试
	}
	t.Log("TestSub() 测试通过")
	t.Logf("期望值：%d, 实际值：%d", -1, Sub(1, 2))
}

// 如果测试用例过多，可以使用类似表格的测试
func TestAdd2(t *testing.T) {
	cases := []struct {
		name         string
		a, b, expect int
	}{
		// 也可以在文件里读取
		{"a1", 1, 2, 3},
		{"a2", 2, -1, 1},
		{"a3", -1, -4, -5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if Add(c.a, c.b) != c.expect {
				t.Errorf("期望值：%d, 实际值：%d", c.expect, Add(c.a, c.b))
			}
			t.Logf("期望值：%d, 实际值：%d", c.expect, Add(c.a, c.b))
		})
	}
}
