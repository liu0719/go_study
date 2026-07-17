package main

import "fmt"

// chan
func Push(c chan string,str string) {
	c<-str
	fmt.Println(str,"已加入chan")
}
func main() {
	var c chan string
	c = make(chan string, 1)
	c <- 
}
