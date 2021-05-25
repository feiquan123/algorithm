// 数组、字符串、切片 转换
package main

/*
#include <string.h>
char arr[10]={67,63,74,75,86,87,38,39,69};
char *s = "Hello";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 通过 reflect.SliceHeader 转化 array
	var arr0 []byte
	var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Hdr.Data = uintptr(unsafe.Pointer(&C.arr[0]))
	arr0Hdr.Len = 10
	arr0Hdr.Cap = 10
	fmt.Println(arr0)

	// 通过切片语法转化 array
	arr1 := (*[31]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]
	fmt.Println(arr1)

	// 通过 reflect.SliceHeader 转化 string
	var s0 string
	var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(C.s))
	s0Hdr.Len = int(C.strlen(C.s))
	fmt.Println(s0)

	// 通过切片语法转化 string
	sLen := int(C.strlen(C.s))
	s1 := string((*[31]byte)(unsafe.Pointer(C.s))[:sLen:sLen])
	fmt.Println(s1)

	// 不同数组类型转化
	var x = []int{1, 2, 3, 4, 5}
	var y = []int64{0}

	xHdr := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	yHdr := (*reflect.SliceHeader)(unsafe.Pointer(&y))

	yHdr.Data = xHdr.Data
	yHdr.Cap = xHdr.Cap * int(unsafe.Sizeof(x[0])) / int(unsafe.Sizeof(y[0]))
	yHdr.Len = xHdr.Len * int(unsafe.Sizeof(x[0])) / int(unsafe.Sizeof(y[0]))
	fmt.Println(y)
}
