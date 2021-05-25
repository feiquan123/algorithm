package main

/*
# include <stdint.h>

union B1{
	int i;
	float f;
};

union B2{
	int8_t i8;
	int64_t i64;
};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var b1 C.union_B1
	fmt.Printf("%T  %v\n", b1, b1) // [4]uint8
	// 通过强制类型转化，解析 union_strut
	fmt.Println("b1.i: ", *(*C.int)(unsafe.Pointer(&b1)))
	fmt.Println("b1.f: ", *(*C.float)(unsafe.Pointer(&b1)))

	var b2 C.union_B2
	fmt.Printf("%T  %v\n", b2, b2) // [8]uint8
}
