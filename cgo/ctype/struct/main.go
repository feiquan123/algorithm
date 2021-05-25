package main

/*
struct A {
	int i;
	float f;
	int type;
	float _type; // 屏蔽CGO对type成员的访问
	int size:10; // 位字段无法访问
	float arr[]; // 零长的数组也无法访问
};
*/
import "C"
import "fmt"

func main(){
	var a C.struct_A
	fmt.Printf("%T  %v\n", a, a)
	fmt.Println(a.i)
	fmt.Println(a.f)
	fmt.Printf("%T %v\n",a._type,a._type)
	// fmt.Println(a.size)
	// fmt.Println(a.arr)
}