/* Code generated by cmd/cgo; DO NOT EDIT. */

#include <stdlib.h>
#include "_cgo_export.h"

#pragma GCC diagnostic ignored "-Wunknown-pragmas"
#pragma GCC diagnostic ignored "-Wpragmas"
#pragma GCC diagnostic ignored "-Waddress-of-packed-member"
extern void crosscall2(void (*fn)(void *), void *, int, __SIZE_TYPE__);
extern __SIZE_TYPE__ _cgo_wait_runtime_init_done(void);
extern void _cgo_release_context(__SIZE_TYPE__);

extern char* _cgo_topofstack(void);
#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()


#define _cgo_msan_write(addr, sz)

extern void _cgoexp_ff8dc21411a5__cgo_qsort_compare(void *);

CGO_NO_SANITIZE_THREAD
int _cgo_qsort_compare(void* a, void* b)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	typedef struct {
		void* p0;
		void* p1;
		int r0;
	} __attribute__((__packed__)) _cgo_argtype;
	static _cgo_argtype _cgo_zero;
	_cgo_argtype _cgo_a = _cgo_zero;
	_cgo_a.p0 = a;
	_cgo_a.p1 = b;
	_cgo_tsan_release();
	crosscall2(_cgoexp_ff8dc21411a5__cgo_qsort_compare, &_cgo_a, 20, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return _cgo_a.r0;
}
