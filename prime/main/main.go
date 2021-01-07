package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/feiquan123/algorithm/prime"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: prime Number")
		os.Exit(1)
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	for _, v := range prime.NPrimeEratosthenes(n) {
		fmt.Println(v)
	}
}
