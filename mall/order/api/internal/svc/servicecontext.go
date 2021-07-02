package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"zero/mall/order/api/internal/config"
	"zero/mall/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
