package main

import (
	"fmt"
	//"time"
)

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	fmt.Println("hello")
	return ch
}
func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	fmt.Println("nice work!!")
}
func main() {
	suck(pump())
}
