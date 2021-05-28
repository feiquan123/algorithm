package main

import "C"
import "unsafe"

//export NewGoString
func NewGoString(s *C.char) *C.char {
	gs := C.GoString(s)
	id := NewObjectID(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := ObjectID(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoString
func PrintGoString(p *C.char) {
	id := ObjectID(uintptr(unsafe.Pointer(p)))
	gs := id.Get().(string)
	print(gs)
}
