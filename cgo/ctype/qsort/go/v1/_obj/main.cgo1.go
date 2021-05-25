// Code generated by cmd/cgo; DO NOT EDIT.

//line /Users/zhi/Program/GoProgram/src/github.com/feiquan123/algorithm/cgo/ctype/qsort/go/v1/main.go:1:1
package main

// extern int go_qsort_compare(void* a, void *b);
import _ "unsafe"
import (
	"fmt"
	"unsafe"

	"github.com/feiquan123/algorithm/cgo/ctype/qsort/go/v1/qsort"
)

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer)  /*line :13:44*/_Ctype_int /*line :13:49*/ {
	pa, pb := (* /*line :14:14*/_Ctype_int /*line :14:19*/)(a), (* /*line :14:27*/_Ctype_int /*line :14:32*/)(b)
	return  /*line :15:9*/_Ctype_int /*line :15:14*/(*pa - *pb)
}

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}
	qsort.Sort(
		unsafe.Pointer(&values[0]),
		len(values),
		int(unsafe.Sizeof(values[0])),
		qsort.CompareFunc(( /*line :24:21*/_Cgo_ptr(_Cfpvar_fp_go_qsort_compare) /*line :24:38*/)),
	)
	fmt.Println(values)
}
