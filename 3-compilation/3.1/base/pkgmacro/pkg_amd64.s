
#define SWAP(x,y,t) MOVQ x,t; MOVQ y, x; MOVQ t,y

TEXT ·Swap(SB),$0-32
	MOVQ a+0(FP), AX  // AX = a
	MOVQ b+8(FP), BX  // BX = b 

	SWAP(AX,BX,CX)  // AX, BX = BX, AX

	MOVQ AX, ret0+16(FP) 
	MOVQ BX, ret1+24(FP)
	RET
