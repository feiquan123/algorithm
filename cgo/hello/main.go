package main

// #include "hello.h"
import "C"

func main() {
	C.SayHelloC(C.CString("hello world, C"))
	C.SayHelloCPP(C.CString("hello world, C++"))
	C.SayHelloGO(C.CString("hello world, GO"))
}
