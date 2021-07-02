package hanlder

const  HelloServiceName = "handler/HelloService"

type HelloService struct {
}

func (s *HelloService) Hello(name string, reply *string) error {
	// 通过修改reply的值获取到返回值
	*reply = "v2: Hello, " + name
	return nil
}