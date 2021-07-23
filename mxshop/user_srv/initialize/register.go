package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"zero/mxshop-api/user-web/utils"
	"zero/mxshop/user_srv/global"
)

func RegisterService() {

	consulConfig := global.ServerConfig.ConsulConfig
	serverConfig := global.ServerConfig

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic("服务注册失败")
	}

	port, err := utils.GetFreePort(serverConfig.Host)
	if err != nil {
		panic("获取port失败")
	}

	serverConfig.Port = port

	address := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	check := &api.AgentServiceCheck{
		GRPC:                           address,
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}
	registration := api.AgentServiceRegistration{
		Address: serverConfig.Host,
		Port:    serverConfig.Port,
		Name:    serverConfig.Name,
		Check:   check,
		ID:      serverConfig.Name,
		Tags:    []string{"user_srv"},
	}

	err = client.Agent().ServiceRegister(&registration)
	fmt.Println("服务注册******")
	if err != nil {
		panic("服务注册失败" + err.Error())
	}

}
