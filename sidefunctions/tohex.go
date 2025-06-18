package sidefunctions

import "strconv"

func ToHex(s string)string{
	hex,err:= strconv.ParseInt(s,16,64)
if err!=nil {
	return ""
}
return strconv.Itoa(int(hex))
}