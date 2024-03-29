package main

// #include <stdio.h>
import "C"
import "unsafe"

func main() {
	buf := NewMyBuffer(1 << 10)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
