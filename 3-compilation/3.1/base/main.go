package main

import (
	"fmt"

	"github.com/feiquan123/algorithm/3-compilation/3.1/pkgarr"
	"github.com/feiquan123/algorithm/3-compilation/3.1/pkgbool"
	"github.com/feiquan123/algorithm/3-compilation/3.1/pkgfloat"
	"github.com/feiquan123/algorithm/3-compilation/3.1/pkgfunc"
	"github.com/feiquan123/algorithm/3-compilation/3.1/pkgint"
	"github.com/feiquan123/algorithm/3-compilation/3.1/pkgmap"
	"github.com/feiquan123/algorithm/3-compilation/3.1/pkgstring"
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
	pkgfunc.Println()
}
