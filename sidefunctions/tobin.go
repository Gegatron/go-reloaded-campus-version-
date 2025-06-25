package sidefunctions

import "strconv"

func ToBin(s string)string{
	b:=false
	if s[len(s)-1]=='\n' {
		b=true
		s=s[:len(s)-1]
	}
bin,err:= strconv.ParseInt(s,2,64)
if err!=nil {
	return s
}
if b {
	return strconv.Itoa(int(bin))+"\n"
}
return strconv.Itoa(int(bin))
}