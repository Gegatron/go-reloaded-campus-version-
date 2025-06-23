package sidefunctions

import "strconv"

func ToBin(s string)string{
bin,err:= strconv.ParseInt(s,2,64)
if err!=nil {
	return s
}
return strconv.Itoa(int(bin))
}