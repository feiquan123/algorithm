package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// go build -o py3-config py3-config.go
// PKG_CONFIG=./py3-config go build -buildmode=c-shared -o gopkg.so main.go
func main() {
	for _, s := range os.Args {
		if s == "--cflags" {
			out, _ := exec.Command("python3-config", "--cflags").CombinedOutput()
			out = bytes.Replace(out, []byte("-arch"), []byte{}, -1)
			out = bytes.Replace(out, []byte("i386"), []byte{}, -1)
			out = bytes.Replace(out, []byte("x86_64"), []byte{}, -1)
			fmt.Println(string(out))
			return
		}
		if s == "--libs" {
			out, _ := exec.Command("python3-config", "--ldflags").CombinedOutput()
			fmt.Println(string(out))
			return
		}
	}
}
