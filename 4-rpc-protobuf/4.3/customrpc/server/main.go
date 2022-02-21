package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.3/customrpc/api"
)

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return errors.New("auth failed")
	}

	log.Println("login ok")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return errors.New("please login")
	}
	*reply = fmt.Sprintf("hello:%s, from:%s", request, p.conn.RemoteAddr().String())
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.RegisterName(api.HelloServiceName, &HelloService{
				conn: conn,
			})
			p.ServeConn(conn)
		}()
	}
}
