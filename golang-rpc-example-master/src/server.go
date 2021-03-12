package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"server"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	fmt.Print("Press enter key to end server")
	fmt.Scanln()
}
