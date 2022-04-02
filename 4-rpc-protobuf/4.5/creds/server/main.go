package main

import (
	"context"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.5/creds/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *api.String) (*api.String, error) {
	replay := &api.String{
		Value: "hello:" + args.GetValue(),
	}
	return replay, nil
}

func main() {
	testData := filepath.Join(os.Getenv("GOPATH"), "src", "github.com/feiquan123/algorithm/4-rpc-protobuf/4.5/creds/testdata")
	creds, err := credentials.NewServerTLSFromFile(
		filepath.Join(testData, "server.crt"),
		filepath.Join(testData, "server.key"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer(grpc.Creds(creds))

	api.RegisterHelloServiceServer(server, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve(lis)
}
