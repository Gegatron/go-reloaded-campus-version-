package sidefunctions

import "strconv"

func ToBin(s string)string{
bin,err:= strconv.ParseInt(s,2,64)
if err!=nil {
	return ""
}
return strconv.Itoa(int(bin))
}