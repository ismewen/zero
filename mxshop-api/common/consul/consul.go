package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func NewUserWebConsulClient() (*api.Client, error) {
	address := "192.168.101.74:8500"
	cfg := api.DefaultConfig()
	cfg.Address = address
	return api.NewClient(cfg)
}

func Register(address string, port int, name string, id string, tags []string) {
	client, err := NewUserWebConsulClient()
	if err != nil {
		panic(err.Error())
	}

	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.101.74:8001/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registration := api.AgentServiceRegistration{
		Address: address,
		Name:    name,
		ID:      id,
		Port:    port,
		Tags:    tags,
		Check:   check,
	}

	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		panic(err.Error())
	}

}

func AllServices() {
	client, err := NewUserWebConsulClient()
	if err != nil {
		panic("创建client 失败")
	}
	data, err := client.Agent().Services()
	for key, value := range data {
		fmt.Println(key)
		fmt.Println(value)
	}
}

type RpcService struct {
	Host string
	Port string
}

func GetService(serviceId string) (*api.AgentService, error) {
	client, err := NewUserWebConsulClient()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	agentService, _, err := client.Agent().Service(serviceId, nil)
	return agentService, err
}
