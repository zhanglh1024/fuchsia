package main

import (
	"os"
)

func main() {
	err := os.Mkdir("/tmp/test/", 00776)
	if err != nil {
		panic(err)
	}
}
