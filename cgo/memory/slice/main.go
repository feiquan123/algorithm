// package main go 无法创建大于 2G 的内存
// go 调用 c
package main

/*
#include <stdlib.h>

void* makeslice(size_t memsize){
	return malloc(memsize);
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func makeBytesSlice(n int) []byte {
	v := make([]byte, 0)
	pv := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	pv.Data = uintptr(C.makeslice(C.size_t(n)))
	pv.Len = n
	pv.Cap = n
	return v
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func main() {
	s := makeBytesSlice(1<<32 + 1)
	s[len(s)-1] = 125
	fmt.Printf("slice length: %d, last element:%d \n", len(s), s[len(s)-1])
	freeByteSlice(s)
}
