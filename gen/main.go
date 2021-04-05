package main

import "fmt"

type ErrCode int

//go:generate stringer -type ErrCode -linecomment
const (
	ERR_CODE_OK             ErrCode = iota // OK
	ERR_CODE_INVALID_PARAMS                // 无效参数
	ERR_CODE_TIMEOUT                       // 超时
)

func main() {
	fmt.Println(ERR_CODE_OK)
	fmt.Println(ERR_CODE_INVALID_PARAMS)
	fmt.Println(ERR_CODE_TIMEOUT)
}
