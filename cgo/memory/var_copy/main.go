package main

/*
# include <stdlib.h>
# include <stdio.h>

void printString(const char* s){
	printf("%s\n",s);
}
*/
import "C"
import "unsafe"

func printString(s string) {
	cs := C.CString(s) // 内存复制
	defer C.free(unsafe.Pointer(cs))

	C.printString(cs)
}

func main() {
	s := "hello"
	printString(s)
}
