package initialize

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"zero/mxshop/common/utils"
	"zero/mxshop/goods_srv/global"
)

func RegisterServiceSingle() {

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

var ServiceID string = ""
var ConsulClient *api.Client

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
	serviceID := uuid.New().String()
	registration := api.AgentServiceRegistration{
		Address: serverConfig.Host,
		Port:    serverConfig.Port,
		Name:    serverConfig.Name,
		Check:   check,
		ID:      serviceID,
		Tags:    []string{"user_srv"},
	}

	err = client.Agent().ServiceRegister(&registration)
	fmt.Println("服务注册******")
	if err != nil {
		panic("服务注册失败" + err.Error())
	}

	ServiceID = serviceID
	ConsulClient = client
}

func ListenSignal() {
	// 监听退出信号
	consulConfig := global.ServerConfig.ConsulConfig
	serverConfig := global.ServerConfig

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic("服务注销失败")
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = client.Agent().ServiceDeregister(ServiceID)
	if err != nil {
		panic("服务注销失败")
	}
	zap.S().Infof("Name:%s, Service %s, 退出成功", serverConfig.Name, ServiceID)
}
