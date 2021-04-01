// +build debug

package main

import "fmt"

func Debug(v interface{}) {
	fmt.Printf("%#v\n", v)
}
