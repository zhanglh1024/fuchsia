package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Qoutient struct {
	Quo, Rem int
}
type Arinth int

func (arin *Arinth) Muxty(arg *Args, reply *int) error {
	rereply = arg.A * arg.B
	return nil
}
func (t *Arinth) Divide(arg *Args, qui *Qoutient) error {
	if arg.B == 0 {
		return errors.New("divide by zero")
	}
	qui.Quo = arg.A / arg.B
	qui.Rem = arg.A / arg.B
	return nil
}

func main() {
	arinth := new(Arinth)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
