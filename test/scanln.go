package main

import (
	"fmt"
)

func main() {
	testFmt()
}
func testFmt() {
	const l = 3
	arr := make([]string, l, l)
	var word string
	for i := 0; i < l; i++ {
		fmt.Scan(&word)
		arr[i] = word
		fmt.Println()
		word = "word"
	}
	for i := 0; i < l; i++ {
		fmt.Print(arr[i] + " ")
	}
}
