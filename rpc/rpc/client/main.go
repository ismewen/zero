package main

import (
	"fmt"
	"zero/rpc/rpc/client/client_proxy"
)

func main() {
	addr := ":9999"
	protocol := "tcp"

	client := client_proxy.NewHelloServiceClient(protocol, addr)
	var reply string
	err := client.Hello("Ethan", &reply)
	if err != nil {
		fmt.Println("Call Error", err.Error())
		return
	}
	fmt.Println("Call Success", reply)

}
