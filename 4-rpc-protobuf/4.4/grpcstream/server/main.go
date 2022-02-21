package main

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.4/grpcstream/api"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *api.String) (*api.String, error) {
	replay := &api.String{
		Value: "hello:" + args.GetValue(),
	}
	return replay, nil
}

func (p *HelloServiceImpl) Channel(stream api.HelloService_ChannelServer) error {
	for {
		// 接收客户端发送的数据
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &api.String{Value: "hello:" + args.GetValue()}

		// 发送给客户端的响应
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
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
