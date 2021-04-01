package main

import "C"

import "fmt"

//export SayHelloGO
func SayHelloGO(s *C.char) {
	fmt.Println(C.GoString(s))
}
