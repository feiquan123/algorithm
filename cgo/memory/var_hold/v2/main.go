package main

/*
# include <gostring.h>

static void printString(char* s){
	char* gs = NewGoString(s);
	PrintGoString(gs);
	FreeGoString(gs);
}
*/
import "C"

func main() {
	C.printString(C.CString("hello world\n"))
}
