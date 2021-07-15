package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"zero/mxshop/user_srv/handler"
	"zero/mxshop/user_srv/proto"
)

func main() {

	// 启动 grpc
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	PORT := flag.Int("port", 8181, "端口号")
	flag.Parse()
	fmt.Println(IP, PORT)
	server := grpc.NewServer()
	address := fmt.Sprintf("%s:%d", *IP, *PORT)
	fmt.Println(address)
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	err = server.Serve(lis)
}
