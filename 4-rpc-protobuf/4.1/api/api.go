package api

import "net/rpc"

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, replay *string) error
}

func RegisterHelloService(srv *rpc.Server, svc HelloServiceInterface) error {
	return srv.RegisterName(HelloServiceName, svc)
}
