package sidefunctions

import "strconv"

func ToHex(s string)string{
	b:=false
	if s[len(s)-1]=='\n' {
		b=true
		s=s[:len(s)-1]
	}
	hex,err:= strconv.ParseInt(s,16,64)
if err!=nil {
	return "<Not A Hexadecimal Number>"
}
if b {
	return strconv.Itoa(int(hex))+"\n"
}
return strconv.Itoa(int(hex))
}