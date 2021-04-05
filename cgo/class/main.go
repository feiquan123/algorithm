package main

/*
#include <stdio.h>

const char* cs_main = "hello world";

void printint(int v){
	printf("print int:%d\n",v);
}
*/
import "C"
import (
	"fmt"

	"github.com/feiquan123/algorithm/cgo/class/help"
)

var CS = C.cs_main

func main() {
	C.printint(C.int(10))

	// fmt.Println(CS==help.CS) // 类型不一致
	// help.PrintCString(C.cs)  // 类型不一致
	// C.puts(help.CS) // 类型不一致
	C.puts(CS)

	// class
	c := help.New("hello world")
	c.Println()
	fmt.Println("GoString: ", c)
	fmt.Printf("%#v\n", CS)
	fmt.Printf("%#v\n", help.CS)
}
