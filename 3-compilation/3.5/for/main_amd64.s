

TEXT ·LoopAdd(SB),$0-32
	MOVQ cnt+8*0(FP), AX // cnt
	MOVQ v0+8*1(FP), BX // v0 /result
	MOVQ step+8*2(FP), CX // step

	MOVQ BX, R8 // next
	MOVQ $1, DX // i

LOOP_IF:
	CMPQ DX, AX // compare i, cnt
	JL LOOP_BODY // if i < cnt; goto LOOP_BODY
	JMP LOOP_END

LOOP_BODY:
	ADDQ CX, R8 // next += step
	ADDQ R8, BX // result += next
	ADDQ $1, DX // i++
	JMP LOOP_IF

LOOP_END:
	MOVQ BX, ret+8*3(FP) // return result
	RET
