package main

import (
	"fmt"

	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgarr"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgbool"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgfloat"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgfunc"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgfunc/next"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgint"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgmacro"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgmap"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgstring"
	"github.com/feiquan123/algorithm/3-compilation/3.1/base/pkgstruct"
)

func main() {
	fmt.Println("bool\t", pkgbool.BoolValue, pkgbool.TrueValue, pkgbool.FalseValue)
	fmt.Println("int\t", pkgint.Id, pkgint.Count)
	fmt.Println("float\t", pkgfloat.Float32Value, pkgfloat.Float64Value)
	fmt.Println("string\t", pkgstring.Name, pkgstring.Name24)
	fmt.Println("arr\t", pkgarr.Num)
	fmt.Println("slice\t", string(pkgstring.NameSlice))
	fmt.Println("map\t", pkgmap.M)
	fmt.Println("channel\t", pkgmap.C)
	fmt.Printf("struct\t%#v\n", pkgstruct.P)
	pkgfunc.Println()
	a, b := pkgfunc.Swap(1, 2)
	fmt.Println("func args swap\t", a, b)
	// fmt.Println("func alloc memory\t", string(pkgfunc.Foo(true, 10)))
	a, b = pkgmacro.Swap(10, 20)
	fmt.Println("macro func swap\t", a, b)

	// 函数流程控制
	next.Calc()
}
