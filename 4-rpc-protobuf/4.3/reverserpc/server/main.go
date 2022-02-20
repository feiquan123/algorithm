package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.1/api"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "hello:" + request
	return nil
}

func main() {
	err := api.RegisterHelloService(rpc.DefaultServer, new(HelloService))
	if err != nil {
		log.Fatal("register error:", err)
	}

	for {
		fmt.Println("----------")
		conn, err := net.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Println("dialing error:", err)
			time.Sleep(time.Second)
			continue
		}

		if conn == nil {
			log.Println("conn is nil")
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}
