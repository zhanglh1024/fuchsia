package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	select {
	case <-time.After(time.Second * 2):
		break
	case c <- i:
		println("i")
	}
	fmt.Println("+++++")
	for {
		fmt.Println("++++")
		select {
		case <-time.After(time.Second * 4):
			fmt.Println("time out!!!")
			break
		case j := <-c:
			fmt.Println(j)
		}
		fmt.Println("++++++")
	}
	fmt.Println("---------")
}
