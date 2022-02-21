package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.4/grpcstream/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal("channel error:", err)
	}
	go func() {
		// 给服务端发送数据
		for {
			if err := stream.Send(&api.String{
				Value: "world",
			}); err != nil {
				log.Fatal("send error:", err)
			}
			time.Sleep(time.Second)
		}
	}()

	// 接收服务端的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
