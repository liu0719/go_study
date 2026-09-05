package main

import (
	"fmt"
	"sync"
)

// 线程安全

// 普通map多协程操作会报错
var maps = map[int]string{}

// 阻塞
var wait = sync.WaitGroup{}

// 互斥
var lock = sync.Mutex{}

func 加(num *int) {
	lock.Lock()
	for i := 0; i < 10000; i++ {

		*num++

	}
	lock.Unlock()
	wait.Done()
}
func 减(num *int) {
	lock.Lock()
	for i := 0; i < 10000; i++ {

		*num--

	}
	lock.Unlock()
	wait.Done()
}
func main() {

	var mapm = sync.Map{}
	num := 0
	wait.Add(2)
	go 加(&num)
	go 减(&num)
	wait.Wait()
	fmt.Println(num)

	// 布尔值,可以查看当前是上锁了还是没上锁
	fmt.Println(lock.TryLock())

	// 协程map

	// concurrent map read and map write
	go func() {
		for {
			// maps[1] = "张三"
			// 类似精简指令集,用load和store来存取
			mapm.Store(1, "张三")
			// 循环map
			// mapm.Range()
			// 取到且删除
			// mapm.LoadAndDelete()
			// 有就取,没有就创建
			// mapm.LoadOrStore()

		}
	}()

	// 会卡住主程序
	// select {}
	for {
		// fmt.Println(maps[1])
		fmt.Println(mapm.Load(1))
	}

}
