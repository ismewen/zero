package main

import (
	"fmt"
	"net"
	"net/rpc"
	"zero/rpc/rpc/hanlder"

	"zero/rpc/rpc/server/server_proxy"
)

func main() {
	// 实例化一个server
	addr := ":9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic("启动失败")
	}
	// 注册服务
	_ = server_proxy.RegisterHelloService(&hanlder.HelloService{})
	// 启动rpc server
	fmt.Println("listening on ", addr)
	for {
		conn, _ := listener.Accept() // 接收一个连接
		go rpc.ServeConn(conn)       // 转交给rpc处理
	}
}
