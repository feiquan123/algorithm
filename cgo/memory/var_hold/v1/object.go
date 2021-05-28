package main

import "sync"

type ObjectID int32

// 对象存储，id 和 go 对象绑定
var refs struct {
	sync.Mutex
	objs map[ObjectID]interface{}
	next ObjectID
}

// 初始化创建全局对象
func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[ObjectID]interface{})
	refs.next = 1 << 10
}

// NewObjectID 绑定 id 和 go 对象，返回 id
func NewObjectID(obj interface{}) ObjectID {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++
	refs.objs[id] = obj
	return id
}

// IsNil 判断当前 id 是否是无效的
func (id ObjectID) IsNil() bool {
	return id == 0
}

// Get 通过 id 获取绑定的 go 对象
func (id ObjectID) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[id]
}

// Free 通过 id 解绑 go 对象，释放内存
func (id *ObjectID) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[*id]
	delete(refs.objs, *id)
	return obj
}
