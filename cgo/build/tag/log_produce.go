// +build !debug

package main

import "fmt"

func Debug(v interface{}) {
	fmt.Println("produce: ",v)
}
