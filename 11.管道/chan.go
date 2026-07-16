package main

import "fmt"

// chan

func main() {
	var c chan int
	c = make(chan int, 1)
	c <- -8
	fmt.Println(<-c)
}
