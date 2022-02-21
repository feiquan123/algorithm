package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.3/customrpc/api"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ api.HelloServiceInterface = (*HelloServiceClient)(nil)

func (p *HelloServiceClient) Login(request string, reply *string) error {
	return p.Client.Call(api.HelloServiceName+".Login", request, reply)
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(api.HelloServiceName+".Hello", request, reply)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalln("dialing error:", err)
	}
	client := &HelloServiceClient{
		Client: rpc.NewClient(conn),
	}

	var replay string
	err = client.Login("user:password", &replay)
	if err != nil {
		log.Fatalln("Login error:", err)
	}
	err = client.Hello("world", &replay)
	if err != nil {
		log.Fatalln("Hello error:", err)
	}
	fmt.Println(replay)
}
