package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/feiquan123/algorithm/pubsub"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	allSub := p.Subscribe()
	golangSub := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	// send msg
	go func() {
		i := 0
		for {
			msg := strconv.Itoa(i) + "\thello, "
			if i%2 == 0 {
				msg += "all"
			} else {
				msg += "golang"
			}
			p.Publish(msg)
			i++

			time.Sleep(500 * time.Millisecond)
		}
	}()

	// consumer msg
	go func() {
		for msg := range allSub {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golangSub {
			fmt.Println("goalng:", msg)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM)
	fmt.Println("signal: ", <-sig)
}
