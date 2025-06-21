package sidefunctions

import "strings"

func Punc(str []string) []string {

	for i := 0; i < len(str); i++ {
		if i != 0 && !NotPunc(str[i]) {
			for j := i-1; j >= 0; j-- {
				if NotType(str[j]) {
					str[j] = str[j] + str[i]
					str[i] = ""
					str = Clean(SliceToString(str))
					i--
					break
				}
			}
		
		}

	}
	return str
}
func IsPunc(s byte) bool {
	if s == ',' || s == '.' || s == ':' || s == ';' || s == '?' || s == '!' {
		return true

	}
	return false
}
func NotType(s string) bool {
	if s == "(up)" || s == "(low)" || s == "(cap)" || s == "(hex)" || s == "(bin)" || strings.Contains(s,"(cap,")|| strings.Contains(s,"(up,")|| strings.Contains(s,"(low,"){
		return false
	}
	return true
}