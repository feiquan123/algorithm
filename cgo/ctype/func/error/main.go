package main

/*
#include <errno.h>

int div(int a, int b) {
	if (b==0){
		errno = EINVAL;
		return 0;
	}
	return a/b;
}

void noreturn(){};
*/
import "C"
import "fmt"

func main() {
	v0, err0 := C.div(2, 1)
	fmt.Println(v0, err0)

	v1, err1 := C.div(2, 0)
	fmt.Println(v1, err1)

	v2, err2 := C.noreturn()
	fmt.Printf("%T\n", v2)
	fmt.Println(v2, err2)
}