/*Package soundex algorith:
1. 保存姓名的首字母
2. 英文字符 codes 规则替换, 为 “0” 时删除
3. 在原文字中，连续出现相同的字符只保留第一次出现者，其余删除
4. 形成 “字母，数字，数字，数字” 的形式，只保留3位，不足补 “0”
*/
package soundex

import (
	"strings"
)

const (
	code0 = "0" // aehiouwy
	code1 = "1" // bfpv
	code2 = "2" // cgjkqsxz
	code3 = "3" // dt
	code4 = "4" // l
	code5 = "5" // mn
	code6 = "6" // r
)

var (
	// 					 A 		B 	  C    	 D 		E 		F 	  G 	  H 	I 		J 	  K 	L 		M 		N 	  O 	  P 	Q 		R 		S 	  T 	U   	V 	  W 	  X 	 Y   	Z
	codes = [26]string{code0, code1, code2, code3, code0, code1, code2, code0, code0, code2, code2, code4, code5, code5, code0, code1, code2, code6, code2, code3, code0, code1, code0, code2, code0, code2}
)

// Soundex soundex ASCII string -> A000
func Soundex(s string) (r string) {
	if len(s) == 0 {
		return ""
	}
	s = strings.ToUpper(s)
	bs := []byte(s)
	rs := make([]byte, 0)
	j := 0
	exists := map[byte]struct{}{}
	for _, b := range bs {
		if b >= 'A' && b <= 'Z' {
			if j == 0 {
				rs = append(rs, b)
				j++
				continue
			}
			if codes[b-'A'] == code0 {
				continue
			}

			if _, ok := exists[b]; ok {
				continue
			} else {
				rs = append(rs, []byte(codes[b-'A'])[0])
				exists[b] = struct{}{}
				j++
				if j > 3 {
					break
				}
			}
		}
	}
	if len(rs) < 4 {
		rs = append(rs, []byte(strings.Repeat("0", 4-len(rs)))...)
	}
	return string(rs)
}
