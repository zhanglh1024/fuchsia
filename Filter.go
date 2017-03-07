package main

import (
	"fmt"
)

func add(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}
func Fileter(in, out chan int, pum int) {
	for {
		i := <-in
		if i%pum != 0 {
			out <- i
		}
	}
}
func main() {
	ch := make(chan int)
	go add(ch)
	for {
		pum := <-ch
		fmt.Println(" ", pum, "  ")
		ch1 := make(chan int)
		go Fileter(ch, ch1, pum)
		ch = ch1
	}
}
