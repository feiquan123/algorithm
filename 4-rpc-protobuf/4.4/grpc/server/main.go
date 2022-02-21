package main

import (
	"context"
	"log"
	"net"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.4/grpc/api"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *api.String) (*api.String, error) {
	replay := &api.String{
		Value: "hello:" + args.GetValue(),
	}
	return replay, nil
}

func main() {
	grpcServer := grpc.NewServer()
	api.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	grpcServer.Serve(lis)
}
