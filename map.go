package main

import "fmt"

func main() {
	var spChan map[int]chan int
	spChan = make(map[int]chan int)
	for i := 0; i < 8; i++ {
		spChan[i] = make(chan int, 8)
		for j := 0; j < 6; j++ {
			spChan[i] <- j
		}
	}
	fmt.Println(len(spChan))
}
