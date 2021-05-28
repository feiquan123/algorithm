package main

/*
# include <stdio.h>

void printString(const char*s , int n){
	int i;
	for (i=0; i<n; i++){
		putchar(s[i]);
	}
	putchar('\n');
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func printString(s string) {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	// c 函数调用期间， 这个 p.Data 的内存地址不变
	// 这个 goroutine 阻塞，不会发生扩容操作
	C.printString((*C.char)(unsafe.Pointer(p.Data)), C.int(len(s)))
}

func main() {
	s := "hello world"
	printString(s)
}
