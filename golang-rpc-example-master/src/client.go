package main

import (
	"fmt"
	"log"
	"net/rpc"
	"server"
)

func main() {
	fmt.Println("Hello, world!")

	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 7 * 8
	var reply int
	args := new(server.Args)
	args.A = 7
	args.B = 8

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:d", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// 15 / 3
	args.A = 15
	args.B = 3
	quot := new(server.Quotient)

	err = client.Call("Arith.Divide", args, quot)
	if err != nil {
		log.Fatal("arith error:d", err)
	}
	fmt.Printf("Arith: %d/%d= quotient:%d remainder:%d\n", args.A, args.B, quot.Quo, quot.Rem)
}
