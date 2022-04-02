package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.4/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//go:generate env GODEBUG="x509ignoreCN=0" go run . 
func main() {
	testData := filepath.Join(os.Getenv("GOPATH"), "src", "github.com/feiquan123/algorithm/4-rpc-protobuf/4.5/creds/testdata")
	creds, err := credentials.NewClientTLSFromFile(
		filepath.Join(testData, "server.crt"),
		"server.grpc.io", // CA 证书生成时指定的服务器名
	)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := api.NewHelloServiceClient(conn)
	replay, err := client.Hello(context.Background(), &api.String{Value: "world"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(replay.GetValue())
}
