package client_proxy

import (
	"net/rpc"

	"zero/rpc/rpc/hanlder"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string, )HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connect error")
	}
	return HelloServiceStub{conn}
}

func (stub *HelloServiceStub )Hello(request string, reply *string) error {
	err := stub.Call(hanlder.HelloServiceName + ".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}