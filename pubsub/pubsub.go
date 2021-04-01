// Package pubsub implements a simple multi-topic pub-sub library
package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // 订阅者是一个通道
	topicFunc  func(v interface{}) bool // 主题为一个过滤器
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex
	buffer      int                      // 订阅队列的缓存大小
	timeout     time.Duration            // 发布者超时时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// 构建一个发布者对象, 可以设置超时时间和缓存队列的长度
func NewPublisher(publicTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publicTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者, 订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) subscriber {
	ch := make(subscriber, p.buffer)
	p.m.Lock()
	defer p.m.Unlock()

	p.subscribers[ch] = topic
	return ch
}

// 添加一个新的订阅者, 订阅全部主题
func (p *Publisher) Subscribe() subscriber {
	return p.SubscribeTopic(nil)
}

// 退出订阅
func (p *Publisher) Evict(sub subscriber) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发送主题，可以容忍一定程度的超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// 发布主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭发布者对象，同时关闭所有的订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}
