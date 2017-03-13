package main

import (
	"fmt"
	"time"
)

func add1(ch chan int) {
	for i := 0; i < 6; i++ {
		ch <- i
	}
}//hello nice work
func add2(ch chan int) {
	for i := 0; i < 6; i++ {
		ch <- i
	}
}
func outa(ch1, ch2 chan int) {
	for i := 0; ; i++ {
		select {
		case v := <-ch1:
			fmt.Println("channel ch1 out", v)
		case v := <-ch2:
			fmt.Println("channel ch2 out", v)
		}
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go add1(ch1)
	go add2(ch2)
	go outa(ch1, ch2)
	time.Sleep(2 * 1e9)
}
