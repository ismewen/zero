package utils

import (
	"fmt"
	"net"
)

func GetFreePort(host string) (int, error) {
	// 理论上应该加全局的锁
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:0", host))
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)

	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil

}
