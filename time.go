package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix() - 29*24*3600
	tm := time.Unix(timestamp, 0)
	time := tm.Format("2006-01-03 15:04:05")
	fmt.Println(time)
}
