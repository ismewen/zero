package server_proxy

import (
	"net/rpc"

	"zero/rpc/rpc/hanlder"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}


func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}

