package main

import (
	"unsafe"
)

type MyBuffer struct {
	cptr *cgo_MyBuffer_T
}

func NewMyBuffer(size int) *MyBuffer {
	return &MyBuffer{
		cptr: cgo_NewMyBuffer(size),
	}
}

func (p *MyBuffer) Delete() {
	cgo_DeleteMyBufffer(p.cptr)
}

func (p *MyBuffer) Data() []byte {
	data := cgo_MyBuffer_Data(p.cptr)
	size := cgo_MyBuffer_Size(p.cptr)
	return ((*[1 << 31]byte)(unsafe.Pointer(data)))[0:int(size):int(size)]
}
