package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublisher bool
	Price       float32
}

func main() {
	gobook := Book{
		"Go programing",
		[]string{"zhanglh", "Deire", "tianaishang"},
		"is zhanglh.com.cn",
		true,
		99.8,
	}
	b, err := json.Marshal(gobook)
	if err == nil {
		fmt.Println(b)
	}
	var book Book
	json.Unmarshal(b, &book)
	fmt.Println(book)
}
