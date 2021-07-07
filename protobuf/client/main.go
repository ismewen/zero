package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"zero/protobuf"
)

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		panic("Connect Error")
	}
	defer conn.Close()

	client := protobuf.NewGreeterClient(conn)
	r, err := client.SayHello(context.Background(), &protobuf.HelloRequest{Name: "Ethan"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r.Message)
}
