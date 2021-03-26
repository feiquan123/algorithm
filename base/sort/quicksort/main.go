package main

import "fmt"

// QuickSort is the quicksort
func QuickSort(arr []int, l, r int) {
	if l < r {
		i, m, j := l, (l+r)>>1, r
		mv := arr[m]
		for i <= j {
			for arr[j] > mv {
				j--
			}
			for arr[i] < mv {
				i++
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if j > l {
			QuickSort(arr, l, j)
		}
		if i < r {
			QuickSort(arr, i, r)
		}
	}
}

func main() {
	arr := []int{0, 1, 3, 2, 7, 8, 2, 6, 9, 3, 11, 0, 8}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
