// panic: runtime error: cgo result has Go pointer
// GODEBUG=cgocheck=0 go run .   取消读返回指针的检查
// GODEBUG=cgocheck=1 go run .   对返回指针简单检查
// GODEBUG=cgocheck=2 go run .   对返回指针完整检查
package main

/*
# include <stdio.h>

extern int* getGoPtr();

static void Main(){
	int* p = getGoPtr();
	*p = 42;
	printf("%d\n",*p);
}
*/
import "C"

func main() {
	C.Main()
}

// 导出 c 函数不能返回 go 内存
//export getGoPtr
func getGoPtr() *C.int {
	return new(C.int)
}
