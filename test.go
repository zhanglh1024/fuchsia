package main

import (
	"fmt"
	"time"
)

func pump(msg string, c chan string) {
	for {
		time.Sleep(1e9)
		c <- msg
	}
}

func main() {
	ch1, ch2 := make(chan string), make(chan string)
	go pump("get the first msg", ch1)
	go pump("get the second mst", ch2)
	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Second):
			fmt.Println("Time Out!!")
		}
	}
	var intput string
	fmt.Scanln(&intput)
}
