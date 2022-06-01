package pkgfunc

import "fmt"

var helloworld string

func output(s string) {
	fmt.Println("func\t", s)
}

func Println()

//go:nosplit
func Swap(a, b int) (int, int)

//go:nosplit
func Foo(a bool, b int16) (c []byte) // 参数和返回值的内存布局

//go:nosplit
func FooLocal() { // 函数中的局部变量
	var c []byte
	var b int16
	var a bool
	_, _, _ = c, b, a
}
