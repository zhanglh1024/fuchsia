package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	c <- "ping"
	c <- "nice"
}
func pinger1(c1, c2 chan string) {
	msg := <-c1
	fmt.Println("Hello Mychannel", msg, c1, c2)
	c2 <- "pong"
}
func printer(c2 chan string) {
	fmt.Println(<-c2)
}

func main() {
	c1, c2 := make(chan string), make(chan string)
	go pinger(c1)
	fmt.Println(<-c1)
	c3 := make(chan string)
	c3 <- "hello world!!"
	go pinger1(c1, c2)
	go printer(c2)
	time.Sleep(4 * 1e9)
	fmt.Println("the main thread over to fast!!")
}
