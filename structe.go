package main

import (
	"fmt"
)

type (
	A struct {
		v int
	}
	B struct {
		A
	}
	C struct {
		*A
	}
)

func (a *A) setV(v int) {
	a.v = v
}
func (a A) getV() int {
	return a.v
}
func (b B) getV() string {
	return "B"
}
func (c *C) getV() bool {
	return true
}
func main() {
	a := A{}
	b := B{}
	c := C{&A{}}
	fmt.Println(a.v)
	fmt.Println(b.v)
	fmt.Println(c.v)
	a.setV(3)
	b.setV(7)
	c.setV(9)
	fmt.Println(a.getV(), b.A.getV(), c.A.getV)
	fmt.Println(a.getV(), b.getV(), c.getV())

}
