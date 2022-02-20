package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

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

func mainSync() {
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

func mainAsync() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 异步调用
	helloCall := client.Go(api.HelloServiceName+".Hello", "world", new(string), nil)

	// 处理其他工作
	time.Sleep(time.Second)
	fmt.Println("处理其他工作 Done")

	// 获取调用结果
	helloCall = <-helloCall.Done
	if err = helloCall.Error; err != nil {
		log.Fatal(err)
	}

	args := helloCall.Args.(string)
	replay := helloCall.Reply.(*string)
	fmt.Println("args:", args)
	fmt.Println("replay:", *replay)
}

func main() {
	// mainSync()
	mainAsync()
}
