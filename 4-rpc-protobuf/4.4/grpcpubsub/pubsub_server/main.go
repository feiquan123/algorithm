package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.4/grpcpubsub/api"
	"google.golang.org/grpc"
)

type PubSubService struct {
	pub *pubsub.Publisher
}

func NewPubSubService() *PubSubService {
	return &PubSubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubSubService) Publish(ctx context.Context, arg *api.String) (*api.String, error) {
	p.pub.Publish(arg.GetValue())
	return &api.String{}, nil
}

func (p *PubSubService) Subscribe(arg *api.String, stream api.PubSubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&api.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	srv := grpc.NewServer()
	api.RegisterPubSubServiceServer(srv, NewPubSubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln(err)
	}
	srv.Serve(lis)
}
