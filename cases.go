package main

import (
	"fmt"
)

func main() {
	ctrl := make(chan bool, 10)
	for i := 0; i < 1000; i++ {
		ctrl <- true
		fmt.Println("----")
		go func() {
			fmt.Println(i)
			defer func() {
				fmt.Println("test", <-ctrl)
			}()
		}()
		fmt.Println("++++++")
	}
}
