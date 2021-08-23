#include "textflag.h"

GLOBL ·M(SB),NOPTR,$8
DATA ·M(SB)/8,$0  // 长度为 0 的 map[string]int

GLOBL ·C(SB),NOPTR,$8
DATA ·C(SB)/8,$0  // 无通道的 int chan
