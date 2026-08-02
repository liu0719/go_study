package main

import (
	"fmt"
	"os"
	"testing"
)

func TestAddition(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("TestAdd() 测试失败")
	}
}

// 构造测试的生命周期
func setup() {
	fmt.Println("setup 执行")
}
func teardown() {
	fmt.Println("teardown 执行")
}

// TestMain 是特殊的测试函数，用于执行包中的其他测试之前执行，函数参数m *testing.M，不是测试用例的参数
func TestMain(m *testing.M) {
	fmt.Println("TestMain 执行")
	setup()
	// 执行所有测试用例,并返回测试结果,1表示失败，0表示成功
	code := m.Run()
	teardown()
	//  将进程的退出状态码传递给操作系统/shell
	// TestMain 必须调用 os.Exit 来传递 m.Run() 的返回值
	// 否则测试结果无法正确传递给外部调用者（IDE、CI 工具等）
	// 可能导致测试失败时仍被判定为成功。
	// exit会立刻退出，导致defer不会被执行，所以不能写defer teardown()，必须在teardown()之后调用os.Exit(code)
	os.Exit(code)
}
