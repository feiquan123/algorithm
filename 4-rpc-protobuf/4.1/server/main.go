package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.1/api"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "hello:" + request
	return nil
}

func rpcMain() {
	err := api.RegisterHelloService(rpc.DefaultServer, new(HelloService))
	if err != nil {
		log.Fatal("register error:", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func httpMain() {
	err := api.RegisterHelloService(rpc.DefaultServer, new(HelloService))
	if err != nil {
		log.Fatal("register error:", err)
	}

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}

func main() {
	// rpcMain()

	/*
		curl -XPOST 'localhost:1234/jsonrpc' -d '{"method":"path/to/pkg.HelloService.Hello","params":["custom msg"],"id":0}'
	*/
	httpMain()
}
