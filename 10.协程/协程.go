package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程

// 默认chan长度为0
// 全局变量的chan可以用来接收各个协程的参数
var moneyChan = make(chan float64)

// 全局变量用来同步协程和主程序

func sing(wait *sync.WaitGroup) {

	fmt.Println("开始唱歌")
	time.Sleep(time.Second * 2)
	fmt.Println("唱歌结束")
	// 结束时wait-1
	defer wait.Done()
}

// 购物函数，把钱写入管道内
func shopping(name string, money float64, wait *sync.WaitGroup) {
	fmt.Println(name, "开始购物")
	time.Sleep(time.Second * 1)
	fmt.Println(name, "购物结束")
	moneyChan <- money
	defer wait.Done()
}

// Goroutine是Go运行时管理的轻量级线程
// 在go中，开启一个协程是非常简单的
func main() {

	wait := sync.WaitGroup{}
	fmt.Println("主线程开始")
	wait.Add(4)
	go sing(&wait)
	go sing(&wait)
	go sing(&wait) 
	go sing(&wait)
	wait.Wait()
	fmt.Println("主线程结束")

	// 添加wait
	wait.Add(3)
	startTime := time.Now()
	go shopping("张三", 2, &wait)
	go shopping("李四", 3, &wait)
	go shopping("王五", 5, &wait)
	// 这里设置一个协程用来处理协程等待，和关闭参数管道
	go func() {
		// 会等到wait为0，才继续执行
		wait.Wait()
		close(moneyChan)
	}()

	// 用来接收从协程里出来的参数
	moneyList := []float64{}
	// 管道里没有参数，for循环会卡在这里等参数
	for money := range moneyChan {
		fmt.Println(money)
		moneyList = append(moneyList, money)

	}
	// 输出
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println(moneyList)

}
