package server

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(name string, reply *string) error {
	// 通过修改reply的值获取到返回值
	*reply = "Hello, " + name
	return nil
}

func main() {
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	addr := ":9999"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic("初始化失败")
	}
}
