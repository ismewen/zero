package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	addr := ":9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic("初始化失败")
	}

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "Ethan", &reply)
	if err != nil {
		panic("call Failed")
	}
	fmt.Println(reply)
}
