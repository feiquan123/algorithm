package api

import "net/rpc"

const KVStoreServiceName = "paht/to/pkg.KVStoreService"

type KVStoreServiceInterface interface {
	Get(key string, value *string) error
	Set(kv [2]string, replay *struct{}) error
	Watch(timeoutSecond int, keyChanged *string) error
}

func RegisterKVStoreService(srv *rpc.Server, svc KVStoreServiceInterface) error {
	return srv.RegisterName(KVStoreServiceName, svc)
}
