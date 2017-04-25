package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func main() {
	stu := &Student{}
	stu.name = "longhu"
	stu.age = 22
	st := CreatStudent()
	fmt.Println(*st == *stu)
	fmt.Println(st)
	fmt.Println(*st)
}
func CreatStudent() *Student {
	stu := &Student{}
	stu.name = "longhu"
	stu.age = 22
	return stu
}
