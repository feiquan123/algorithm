package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"sync"
	"time"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.3/kvstore/api"
)

type KVStoreServiceClient struct {
	*rpc.Client
}

var _ api.KVStoreServiceInterface = (*KVStoreServiceClient)(nil)

func (p *KVStoreServiceClient) Get(key string, value *string) error {
	return p.Client.Call(api.KVStoreServiceName+".Get", key, value)
}

func (p *KVStoreServiceClient) Set(kv [2]string, replay *struct{}) error {
	return p.Client.Call(api.KVStoreServiceName+".Set", kv, replay)
}

func (p *KVStoreServiceClient) Watch(timeoutSecond int, keyChanged *string) error {
	return p.Client.Call(api.KVStoreServiceName+".Watch", timeoutSecond, keyChanged)
}

func DialKVStoreService(network, address string) (*KVStoreServiceClient, error) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		return nil, err
	}
	return &KVStoreServiceClient{
		Client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)),
	}, nil
}

func main() {
	client, err := DialKVStoreService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// watch
		var keyChanged string
		err := client.Watch(10, &keyChanged)
		if err != nil {
			log.Fatal("Watch error:", err)
		}
		fmt.Println("Watch:", keyChanged)
	}()
	time.Sleep(time.Second)

	rand.Seed(time.Now().Unix())
	key := strconv.FormatInt(int64(rand.Int()), 10)
	err = client.Set([2]string{key, key}, new(struct{}))
	if err != nil {
		log.Fatal("Set error:", err)
	}
	wg.Wait()
}
