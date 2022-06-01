package main

import "fmt"

func If(ok, a, b int) int

func main() {
	fmt.Println(If(1, 1, 10))
}

// func If(ok, a, b int) int {
// 	if ok == 0 {
// 		return b
// 	} else {
// 		return a
// 	}
// }

// func If(ok, a, b int) int {
// 	if ok == 0 {
// 		goto L
// 	}
// 	return a
// L:
// 	return b
// }
