package main

import "C"

import (
	"fmt"

	_ "github.com/feiquan123/algorithm/cgo/build/export_mult_go_lib/number"
)

// go build -buildmode=c-archive -o main.a
func main() {
	println("Done")
}

//export goPrintln
func goPrintln(s *C.char) {
	fmt.Println("goPrintln", C.GoString(s))
}
