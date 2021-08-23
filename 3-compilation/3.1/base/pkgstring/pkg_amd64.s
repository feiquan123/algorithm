#include "textflag.h"

GLOBL ·nameData<>(SB),NOPTR,$16
DATA ·nameData<>+0(SB)/16,$"hello world!"

GLOBL ·Name(SB),NOPTR,$16
DATA ·Name+0(SB)/8,$·nameData<>(SB)
DATA ·Name+8(SB)/8,$12

GLOBL ·Name24(SB),NOPTR,$24
DATA ·Name24+0(SB)/8,$·Name24+16(SB)
DATA ·Name24+8(SB)/8,$6
DATA ·Name24+16(SB)/8,$"gopher"

GLOBL ·NameSlice(SB),NOPTR,$40
DATA ·NameSlice+0(SB)/8,$·NameSlice+24(SB)
DATA ·NameSlice+8(SB)/8,$12
DATA ·NameSlice+16(SB)/8,$16
DATA ·NameSlice+24(SB)/12,$"hello world!"
