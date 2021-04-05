package help

// #include <stdio.h>
// const char* cs_help = "hello world";
import "C"

// 导出 C.cs
var CS = C.cs_help

type CChar C.char

func New(s string) *CChar {
	return (*CChar)(C.CString(s))
}

func (p *CChar) String() string {
	return C.GoString((*C.char)(p))
}

func (p *CChar) Println() {
	C.puts(C.CString("CString: " + p.String()))
}

// 参数: *C.cahr => *help.C.cahr, 其他包无法使用
func PrintCString(cs *C.char) {
	C.puts(cs)
}
