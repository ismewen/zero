package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
	"zero/protobuf"
)

type Server struct{}

func (s *Server) SayHello(context.Context, *protobuf.HelloRequest) (*protobuf.HelloReply, error) {
	return &protobuf.HelloReply{Message: "Hello"}, nil
}

func (s *Server) GetStream(req *protobuf.StreamReqData, res protobuf.Greeter_GetStreamServer) error {
	// 服务端流模式, 客户端发送一个请求， 服务端返回一个数据流
	i := 0
	for {
		i++
		time.Sleep(time.Second)
		err := res.Send(&protobuf.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		if err != nil {
			return nil
		}
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *Server) PutStream(stream protobuf.Greeter_PutStreamServer) error {
	// 客户端的流模式， 服务端接受数据流，并返回一个结果
	for {
		resData, err := stream.Recv()
		if err != nil {
			// 正常结束会收到 context canceled 消息
			fmt.Println(err.Error())
			break
		}
		fmt.Println(resData.Data)
	}
	return nil
}
func (s *Server) AllStream(stream protobuf.Greeter_AllStreamServer) error {
	// 服务端的双向流模式
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			// 不断接受来自服务器的数据
			recvData, err := stream.Recv()
			if err != nil {
				return
			}
			fmt.Println("收到客户信息: " + recvData.Data)
		}
	}()

	go func() {
		defer wg.Done()
		// 不断发送 服务器数据
		for {
			time.Sleep(time.Second)
			err := stream.Send(&protobuf.StreamResData{
				Data: "服务端数据",
			})
			if err != nil {
				fmt.Print(err.Error())
				break
			}
		}
	}()
	wg.Wait()
	return nil
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
