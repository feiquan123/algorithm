package main

/*
// 函数签名
int sum(int a, int b);
*/
import "C"

//export sum
func sum(a, b C.int) C.int {
	return a + b
}

func main() {
	println(sum(1,1))
}
