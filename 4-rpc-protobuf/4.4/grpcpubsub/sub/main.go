package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.4/grpcpubsub/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := api.NewPubSubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &api.String{
		Value: "golang",
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		fmt.Println(reply.GetValue())
	}
}
