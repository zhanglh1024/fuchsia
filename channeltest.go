package main

import (
	"fmt"
	"time"
)

func main() {
	chq := sunck()
	pum(chq)
	time.Sleep(1e9)
}
func sunck() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func pum(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
