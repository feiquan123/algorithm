package main

import (
	"fmt"

	"github.com/feiquan123/algorithm/cgo/ctype/qsort/go/v3/qsort"
)

func main() {
	values := []int32{43, 9, 101, 95, 27, 25}

	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
}
