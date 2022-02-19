package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.1/api"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ api.HelloServiceInterface = (*HelloServiceClient)(nil)

func (p *HelloServiceClient) Hello(request string, replay *string) error {
	return p.Client.Call(api.HelloServiceName+".Hello", request, replay)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		Client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)),
	}, nil
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var replay string
	err = client.Hello("custom msg", &replay)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(replay)
}
