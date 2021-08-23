#include "textflag.h"

GLOBL ·BoolValue(SB),RODATA,$1  // 未初始化

GLOBL ·TrueValue(SB),RODATA,$1
DATA ·TrueValue(SB)/1,$1  // 非 0 为true

GLOBL ·FalseValue(SB),RODATA,$1
DATA ·FalseValue(SB)/1,$0
