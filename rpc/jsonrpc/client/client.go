package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddr := os.Args[1]

	client, err := jsonrpc.Dial("tcp", serverAddr) // rpc 转变成 jsonrpc

	if err != nil {
		log.Fatal("dialing err:", err)
	}

	args := Args{12, 2}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith err", err)
	}
	fmt.Printf("arith: %d*%d=%d\n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("arith err", err)
	}
	fmt.Printf("arith: %d/%d=%d\n rem %d", args.A, args.B, quo.Quo, quo.Rem)

}
