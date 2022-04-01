package main

import (
	"context"
	"log"
	"time"

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

	go func() {
		for {
			_, err = client.Publish(context.Background(), &api.String{
				Value: "golang: hello golang, " + time.Now().Format("2016-01-02 15:04:05"),
			})
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			_, err = client.Publish(context.Background(), &api.String{
				Value: "docker: hello docker, " + time.Now().Format("2016-01-02 15:04:05"),
			})
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second)
		}
	}()

	select{}
}
