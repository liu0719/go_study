package main

import (
	"fmt"
	"sync"
)

// chan

// 生产者函数，传入管道和wait
func Push(c chan<- string, str string, wait *sync.WaitGroup) {
	defer wait.Done()
	c <- str
	fmt.Println(str, "已加入chan")
}

func main() {
	c := make(chan string)
	var wait sync.WaitGroup

	// 加入管道的数据
	messages := []string{"你好", "成龙", "特鲁", "小玉", "老爹", "圣主", "瓦龙", "阿福"}
	// wait中加入协程数，用来阻塞
	wait.Add(len(messages))

	// 循环启动协程，加入数据
	for _, msg := range messages {
		go Push(c, msg, &wait)
	}
	// 启动一个协程监听wait耗尽，关闭管道
	go func() {
		wait.Wait()
		close(c)
	}()
	
	// for循环当消费者读出数据，会阻塞住主进程
	ress := []string{}
	for res := range c {
		fmt.Println(res)
		ress = append(ress, res)
	}

	fmt.Println("结果：", ress)
}
