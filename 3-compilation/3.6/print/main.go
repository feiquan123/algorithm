package main

import (
	_ "unsafe"
)

//go:linkname printstring runtime.printstring
func printstring(s string)

//go:linkname printnl runtime.printnl
func printnl()

func printnl_nosplit(s string)

func main() {
	printnl_nosplit("hello world!!!")
}
