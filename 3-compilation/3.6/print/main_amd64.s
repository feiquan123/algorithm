#include "textflag.h"

TEXT ·printnl_nosplit(SB), NOSPLIT, $16-0
	MOVQ s+8*0(FP), AX
	MOVQ AX, 0(SP)

	MOVQ s+8*1(FP), BX
	MOVQ BX, 8(SP)
	CALL runtime·printstring(SB)
	CALL runtime·printnl(SB)
	RET
