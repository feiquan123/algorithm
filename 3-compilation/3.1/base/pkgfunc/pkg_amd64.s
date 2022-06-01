#include "textflag.h"

GLOBL ·helloworld(SB),NOPTR,$40
DATA ·helloworld+0(SB)/8,$·helloworld+16(SB)
DATA ·helloworld+8(SB)/8,$18
DATA ·helloworld+16(SB)/24,$"你好，世界!"

TEXT ·Println(SB),$16-0
	MOVQ ·helloworld+0(SB),AX
	MOVQ AX,0(SP)
	MOVQ ·helloworld+8(SB),BX
	MOVQ BX,8(SP)
	CALL ·output(SB)
	RET

TEXT ·Swap(SB),$0-32
	MOVQ a+0(FP), AX  // AX = a 
	MOVQ b+8(FP), BX  // BX = b
	MOVQ BX, r0+16(FP) // r0 = BX
	MOVQ AX, r1+24(FP) // r1 = AX
	RET

TEXT ·Foo(SB),$0-32
	MOVQ a+0(FP), AX // a
	MOVQ b+2(FP), BX // b
	MOVQ c_dat+8*1(FP), CX // c.Data
	MOVQ c_len+8*2(FP), DX // c.Len
	MOVQ c_cap+8*3(FP), DI // c.Cap
	RET

TEXT ·FooLocal(SB),$32-0
	MOVQ a-32(SP), AX // a
	MOVQ b-30(SP), BX // b
	MOVQ c_data-24(SP), CX // c.Data
	MOVQ c_len-16(SP), DX // c.Len
	MOVQ c_cap-8(SP),DI // c.Cap
	RET