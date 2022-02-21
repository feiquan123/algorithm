package main

import (
	"context"
	"fmt"
	"log"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.4/grpc/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewHelloServiceClient(conn)
	replay, err := client.Hello(context.Background(), &api.String{Value: "world"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(replay.GetValue())
}
