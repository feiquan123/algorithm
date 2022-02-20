package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"time"

	"github.com/feiquan123/algorithm/4-rpc-protobuf/4.3/kvstore/api"
)

type KVStoreService struct {
	m      map[string]string           // 存储键值
	filter map[string]func(key string) // Watch 调用时定义的过滤器函数列表
	mu     sync.Mutex
}

var _ api.KVStoreServiceInterface = (*KVStoreService)(nil)

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	v, ok := p.m[key]
	if !ok {
		return errors.New("not found")
	}
	*value = v
	return nil
}

// replay *struct{} 表示忽略了输出参数
func (p *KVStoreService) Set(kv [2]string, replay *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		// 当修改某个键值对的值时，调用每个过滤器函数
		for _, fn := range p.filter {
			fn(key)
		}
	}
	p.m[key] = value
	fmt.Println(key, value)
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%d-%03d", time.Now().Unix(), rand.Int())
	ch := make(chan string, 10) // buffered

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return errors.New("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}

func main() {
	err := api.RegisterKVStoreService(rpc.DefaultServer, NewKVStoreService())
	if err != nil {
		log.Fatal("register error:", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp error:", err)
	}
	fmt.Println("lister port: 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
