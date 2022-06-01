package main

import "fmt"

func LoopAdd(cnt, v0, step int) int

func main() {
	fmt.Println(LoopAdd(100, 1, 1))
	fmt.Println(LoopAdd(5, 10, -2))
}

// func LoopAdd(cnt, v0, step int) int {
// 	result := v0
// 	next := v0
// 	for i := 1; i < cnt; i++ {
// 		next += step
// 		result += next
// 	}
// 	return result
// }

// func LoopAdd(cnt, v0, step int) int {
// 	var result = v0
// 	var next = v0
// 	var i = 1

// LOOP_IF:
// 	if i < cnt {
// 		goto LOOP_BODY
// 	}
// 	goto LOOP_END

// LOOP_BODY:
// 	next += step
// 	result += next
// 	i++
// 	goto LOOP_IF

// LOOP_END:
// 	return result
// }
