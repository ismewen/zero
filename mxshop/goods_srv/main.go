package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"zero/mxshop/goods_srv/global"
	"zero/mxshop/user_srv/handler"

	//"zero/mxshop/goods_srv/handler"
	"zero/mxshop/goods_srv/initialize"
	"zero/mxshop/user_srv/proto"
)

func main() {

	initialize.InitConfig()
	initialize.InitLogger()

	initialize.RegisterService()

	server := grpc.NewServer()

	address := fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port)
	fmt.Println(address)
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", address)
	// 注册健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	zap.S().Info(address)
	go func() {
		err = server.Serve(lis)
	}()

	// 监听退出信号
	initialize.ListenSignal()
}
