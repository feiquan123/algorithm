package main

import "fmt"

// BinaryFind returns the index of the target
func BinaryFind(arr []int, tag int) int {
	var l, r = 0, len(arr) - 1
	for l <= r {
		m := (l + r) >> 1
		if tag == arr[m] {
			return m
		}
		if tag < arr[m] {
			r = m - 1
			continue
		}
		if tag > arr[m] {
			l = m + 1
			continue
		}
	}
	return -1
}

func main() {
	arr := []int{0, 0, 1, 2, 2, 3, 3, 6, 7, 8, 8, 9, 11}
	fmt.Println(BinaryFind(arr, 2))
	fmt.Println(BinaryFind(arr, 9))
	fmt.Println(BinaryFind(arr, 10))
}
