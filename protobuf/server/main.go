package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"zero/protobuf"
)

type Server struct{}

func (s *Server) SayHello(context.Context, *protobuf.HelloRequest) (*protobuf.HelloReply, error) {
	return &protobuf.HelloReply{Message: "Hello"}, nil
}

func main() {
	g := grpc.NewServer()
	protobuf.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
}
