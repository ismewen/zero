package srv

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"zero/mxshop-api/common/consul"
	"zero/mxshop-api/user-web/global"
	"zero/mxshop-api/user-web/proto"
)

var Srv *ServiceContext

type ServiceContext struct {
	UserRpc proto.UserClient
}

func NewServiceContext() *ServiceContext {
	if Srv != nil {
		return Srv
	}
	// 服务发现
	service, err := consul.GetService(global.ServerConfig.UserSrvConfig.Name)
	if err != nil {
		panic("获取service失败")
	}
	addr := fmt.Sprintf("%s:%d", service.Address, service.Port)
	fmt.Println("xxxxxxxxxxxxxxxxx")
	fmt.Printf(addr)
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
