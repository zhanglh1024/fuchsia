package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}
type Qoutient struct {
	Qui, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(0)
	}
	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{45, 12}
	var reply int
	err = client.Call("Arinth.Muxty", args, &reply)
	if err != nil {
		log.Fatal("arinth error:", err)
	}
	fmt.Printf("Arinth:%d*%d=%d\n", args.A, args.B)
	var qui Qoutient
	err = client.Call("Arinth.Divide", args, &qui)
	if err != nil {
		log.Fatal("arinth error:", err)
	}
	fmt.Printf("Arinth:%d/%d=%d remainder %d\n", args.A, args.B, qui.Qui, qui)
}
