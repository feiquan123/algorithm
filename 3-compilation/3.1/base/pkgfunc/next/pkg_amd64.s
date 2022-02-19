
TEXT ·Calc<>(SB),$24-0
	// 初始化零值
	MOVQ $0, a-8*2(SP) // a = 0
	MOVQ $0, b-8*1(SP) // b = 0

	// 为 a 赋值新变量
	MOVQ $10, AX // AX = 10
	MOVQ AX, a-8*2(SP) // a = AX

	// 以 a 为参数调用打印函数
	MOVQ AX, 0(SP)
	CALL ·printintln(SB)

	// 函数调用后， AX/BX  寄存机可能会被污染，重新赋值
	MOVQ a-8*2(SP), AX // AX = a
	MOVQ b-8*1(SP), BX // BX = b

	// 计算 b 的值, 写入内存
	MOVQ AX, BX
	ADDQ BX, BX
	IMULQ AX, BX  // 结果存入 BX, MULQ 结果存入 AX
	MOVQ BX, b-8*1(SP)

	// 打印 b
	MOVQ BX, 0(SP)
	CALL ·printintln(SB)

	RET

// 简化版
TEXT ·Calc(SB),$16-0
	MOVQ $10, AX // AX = 10
	MOVQ AX, temp-8*1(SP) // temp = AX

	// 打印 temp
	MOVQ AX,0(SP)
	CALL ·printintln(SB)

	// 函数调用后， AX/BX  寄存机可能会被污染，重新赋值
	MOVQ temp-8*1(SP), AX // AX = temp

	MOVQ AX, BX  // BX = AX
	ADDQ BX, BX  // BX += BX
	IMULQ AX, BX // BX *= AX
	
	// 打印 b
	MOVQ BX, 0(SP)
	CALL ·printintln(SB)

	RET



