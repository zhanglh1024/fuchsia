package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}
type Qoutient struct {
	Qui, Rem int
}
type Arinth int

func (arinth *Arinth) Multix(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (arinth *Arinth) Divity(args *Args, qui *Qoutient) error {
	if args.B == 0 {
		return errors.New("the dived can`t be zero")
	}
	qui.Qui = args.A / args.B
	qui.Rem = args.A % args.B
	return nil
}

func main() {
	arinth := new(Arinth)
	rpc.Register(arinth)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
