package main

import (
	"fmt"
	"sync"
	"time"
)

// select

// 默认chan长度为0
// 全局变量的chan可以用来接收各个协程的参数
var moneyChan = make(chan float64)
var nameChan = make(chan string)
var flagChan = make(chan bool)

// 全局变量用来同步协程和主程序

// 购物函数，把钱写入管道内
func send(name string, money float64, wait *sync.WaitGroup) {
	fmt.Println(name, "开始购物")
	time.Sleep(time.Second * 1)
	fmt.Println(name, "购物结束")
	moneyChan <- money
	nameChan <- name
	defer wait.Done()
}

// Goroutine是Go运行时管理的轻量级线程
// 在go中，开启一个协程是非常简单的
func main() {

	wait := sync.WaitGroup{}

	// 添加wait
	wait.Add(3)
	startTime := time.Now()
	go send("张三", 2, &wait)
	go send("李四", 3, &wait)
	go send("王五", 5, &wait)
	// 这里设置一个协程用来处理协程等待，和关闭参数管道
	go func() {
		// 会等到wait为0，才继续执行
		wait.Wait()
		close(moneyChan)
		close(nameChan)
		// 先关闭信号管道，主线程的for会自动等待管道内的值传完才停下，防止读取到无效值，
		close(flagChan)
	}()
	namelist := []string{}
	moneylist := []float64{}
	func() {
		for {
			// select用来异步的从多个channel里面去取数据
			select {
			case name := <-nameChan:
				namelist = append(namelist, name)
				fmt.Println(name, "已从chan中拿出，加到数组")
			case money := <-moneyChan:
				moneylist = append(moneylist, money)
				fmt.Println(money, "已从chan中拿出，加到数组")
			case <-flagChan:
				return
			}

		}
	}()

	// go func() {
	// 	for name := range nameChan {
	// 		fmt.Println(name)
	// 		namelist = append(namelist, name)

	// 	}
	// }()
	// 用来接收从协程里出来的参数

	// 管道里没有参数，for循环会卡在这里等参数
	// for money := range moneyChan {
	// 	fmt.Println(money)
	// 	moneylist = append(moneylist, money)

	// }
	// 输出
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println(moneylist)
	fmt.Println(nameChan)

}
