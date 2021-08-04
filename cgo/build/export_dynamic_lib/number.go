package main

import "C"

// go build -buildmode=c-shared -o number.so
func main() {}

//export number_add_mod
func number_add_mod(a, b, mod C.int) C.int {
	return (a + b) / mod
}
