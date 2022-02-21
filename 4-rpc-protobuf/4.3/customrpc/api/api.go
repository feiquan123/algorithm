package api

import "net/rpc"

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Login(request string, reply *string) error
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv *rpc.Server, svc HelloServiceInterface) error {
	return srv.RegisterName(HelloServiceName, svc)
}
