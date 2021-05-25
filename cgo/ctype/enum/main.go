package main

/*
enum E {
	ZERO,
	ONE,
	TWO,
};
*/
import "C"
import "fmt"

func main() {
	var e C.enum_E = C.TWO
	fmt.Printf("%T  %v\n", e, e)  // uint32

	fmt.Println(C.ZERO)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}
