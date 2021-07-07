package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
	"zero/protobuf"
)

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		panic("Connect Error")
	}
	defer conn.Close()

	client := protobuf.NewGreeterClient(conn)
	// 简单模式

	fmt.Println("简单模式")

	r, err := client.SayHello(context.Background(), &protobuf.HelloRequest{Name: "Ethan"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r.Message)

	// 服务端流模式
	fmt.Println("服务器流模式")

	stream, err := client.GetStream(context.Background(), &protobuf.StreamReqData{Data: "Some Data"})
	if err != nil {
		panic("get stream Failed" + err.Error())
	}
	for {
		a, err := stream.Recv()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(a.Data)
	}
	// 客户端流模式
	fmt.Println("客户端流模式")

	putStream, err := client.PutStream(context.Background())
	if err != nil {
		panic(err.Error())
	}
	i := 0
	for {
		i++
		time.Sleep(time.Second)
		err := putStream.Send(&protobuf.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if i > 10 {
			break
		}
	}

	// 双向模式
	fmt.Println("双向模式")
	wg := sync.WaitGroup{}
	wg.Add(2)
	allStream, err := client.AllStream(context.Background())
	if err != nil {
		panic(err.Error())
	}
	go func() {
		defer wg.Done()
		for {
			recvData, err := allStream.Recv()
			if err != nil {
				return
			}
			fmt.Println("收到服务器数据: %v", recvData.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			err := allStream.Send(&protobuf.StreamReqData{
				Data: "客户端数据",
			})
			if err != nil{
				break
			}
		}
	}()
	wg.Wait()

}
