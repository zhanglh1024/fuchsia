package main

import (
	"fmt"
	//"time"
)

func f(n, m int, ch chan int) {
	for i := 1; i < m; i++ {
		ch <- i * n
	}
	close(ch)
}
func f1(ch chan int, done chan bool) {
	for v := range ch {
		fmt.Println(v) //println test
	}
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go f(3, 6, ch)
	go f1(ch, done)
	fmt.Println(<-done)
	//time.Sleep(2 * time.Second)
}
