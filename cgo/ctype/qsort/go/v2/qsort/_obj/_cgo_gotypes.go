// Code generated by cmd/cgo; DO NOT EDIT.

package qsort

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_int int32

type _Ctype_qsort_cmp_func_t *[0]byte

type _Ctype_size_t = _Ctype_ulong

type _Ctype_ulong uint64

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})
//go:linkname __cgo__cgo_qsort_compare _cgo_qsort_compare
//go:cgo_import_static _cgo_qsort_compare
var __cgo__cgo_qsort_compare byte
var _Cfpvar_fp__cgo_qsort_compare unsafe.Pointer = (unsafe.Pointer)(unsafe.Pointer(&__cgo__cgo_qsort_compare))

//go:cgo_import_static _cgo_4f801c3a363f_Cfunc_qsort
//go:linkname __cgofn__cgo_4f801c3a363f_Cfunc_qsort _cgo_4f801c3a363f_Cfunc_qsort
var __cgofn__cgo_4f801c3a363f_Cfunc_qsort byte
var _cgo_4f801c3a363f_Cfunc_qsort = unsafe.Pointer(&__cgofn__cgo_4f801c3a363f_Cfunc_qsort)

//go:cgo_unsafe_args
func _Cfunc_qsort(p0 unsafe.Pointer, p1 _Ctype_size_t, p2 _Ctype_size_t, p3 *[0]byte) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_4f801c3a363f_Cfunc_qsort, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
		_Cgo_use(p3)
	}
	return
}
//go:cgo_export_dynamic _cgo_qsort_compare
//go:linkname _cgoexp_4f801c3a363f__cgo_qsort_compare _cgoexp_4f801c3a363f__cgo_qsort_compare
//go:cgo_export_static _cgoexp_4f801c3a363f__cgo_qsort_compare
func _cgoexp_4f801c3a363f__cgo_qsort_compare(a *struct {
		p0 unsafe.Pointer
		p1 unsafe.Pointer
		r0 _Ctype_int
	}) {
	a.r0 = _cgo_qsort_compare(a.p0, a.p1)
}
