package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
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

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln("listen error:", err)
	}
	clientChan := make(chan *rpc.Client)

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("accept error:", err)
		}

		clientChan <- rpc.NewClient(conn)
	}()

	doClientWork(clientChan)
}

func doClientWork(clientChan chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	helloServiceClient := &HelloServiceClient{
		Client: client,
	}

	for {
		var replay string
		err := helloServiceClient.Hello("world", &replay)
		if err != nil {
			log.Fatalln("call Hello error:", err)
		}

		fmt.Println(replay)

		time.Sleep(time.Second)
	}
}
