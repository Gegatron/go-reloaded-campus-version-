package sidefunctions

import (
	"strconv"
	
)

func GetNumber(s string) (int,error) {
	new:=""
	b:=false
	for i := 0; i < len(s); i++ {
		if s[i]=='-' && new=="" {
			return strconv.Atoi("0")
		}
		if s[i]>='0' && s[i]<='9' {
			b=true
		}
		if b && !(s[i]>='0' && s[i]<='9') && i!=len(s)-1{
			new=""
			break
		}
		if b && i!=len(s)-1{
			new=new+string(s[i])
		}
	}
	return strconv.Atoi(new)
	
}