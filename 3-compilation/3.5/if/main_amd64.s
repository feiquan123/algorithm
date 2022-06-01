#include "textflag.h"

TEXT ·If(SB),$-32
	MOVQ ok+8*0(FP), CX // ok
	MOVQ a+8*1(FP), AX // a
	MOVQ b+8*2(FP), BX // b

	CMPQ CX, $0 // ok == 0
	JZ L // if ok == 0, goto L
	MOVQ AX, ret+8*3(FP) // return a
	RET

L:
	MOVQ BX, ret+8*3(FP) // return b
	RET

