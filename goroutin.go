package main

import (
	"fmt"
	"time"
)

var i int
var ch = make(chan int)

func main() {
	for range [5]byte{} {
		go Add()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}
func Add() {
	ch <- 0
	i++
	<-ch
}
