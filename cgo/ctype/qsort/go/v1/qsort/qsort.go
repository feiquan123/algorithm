package qsort

/* 
#include <stdlib.h>

typedef int (*qsort_cmp_func_t) (const void* a, const void* b);  // qsort_cmp_func_t 比较函数定义
*/
import "C"
import "unsafe"

type CompareFunc C.qsort_cmp_func_t

// 调用 stdlib.h 中的 qsort
// void	 qsort(void *__base, size_t __nel, size_t __width,
//	int (* _Nonnull __compar)(const void *, const void *));
func Sort(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}