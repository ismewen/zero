package srv

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"zero/mxshop-api/user-web/proto"
)

type ServiceContext struct {
	UserRpc proto.UserClient
}

func NewServiceContext() *ServiceContext {
	addr := "0.0.0.0:8181"
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("UserRpc 连接失败", "error_info", err.Error())
	}
	return &ServiceContext{
		UserRpc: proto.NewUserClient(cc),
	}
}
