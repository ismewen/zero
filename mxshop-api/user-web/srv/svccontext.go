package srv

import (
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important


	"zero/mxshop-api/common/consul"
	"zero/mxshop-api/user-web/global"
	"zero/mxshop-api/user-web/proto"
)

var Srv *ServiceContext

type ServiceContext struct {
	UserRpc proto.UserClient
}

func NewServiceContextBk() *ServiceContext {
	if Srv != nil {
		return Srv
	}
	// 服务发现
	service, err := consul.GetService(global.ServerConfig.UserSrvConfig.Name)
	if err != nil {
		panic("获取service失败")
	}
	addr := fmt.Sprintf("%s:%d", service.Address, service.Port)
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("UserRpc 连接失败", "error_info", err.Error())
	}
	srv := &ServiceContext{
		UserRpc: proto.NewUserClient(cc),
	}
	Srv = srv
	return Srv
}

func NewServiceContext() *ServiceContext {
	consulInfo := global.ServerConfig.ConsulConfig
	userServiceName := global.ServerConfig.UserSrvConfig.Name
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%s/%s?wait=14s", consulInfo.Host, consulInfo.Port, userServiceName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("连接用户服务失败")
	}
	svc := &ServiceContext{
		UserRpc: proto.NewUserClient(conn),
	}
	return svc
}
