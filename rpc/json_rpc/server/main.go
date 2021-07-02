package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(name string, reply *string) error {
	// 通过修改reply的值获取到返回值
	*reply = "Hello, " + name
	return nil
}

func main() {
	// 实例化一个server
	addr := ":9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic("启动失败")
	}
	// 注册服务
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 启动rpc server
	fmt.Println("listening on ", addr)
	for {
		conn, _ := listener.Accept() // 接收一个连接
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
