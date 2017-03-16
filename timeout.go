package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(6e9)
		timeout <- true
	}()
	select {
	case <-ch:
		fmt.Println(time.Now())
	case <-timeout:
		fmt.Println("over!!")
	}

}
