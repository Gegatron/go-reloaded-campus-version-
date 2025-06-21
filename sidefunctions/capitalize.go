package sidefunctions

import "unicode"

func Capitalize(s string)string{
	new:=[]rune(s)
	b:=false
	for i,c:=range new{
		if unicode.IsLetter(c) && !b{
			new[i]=unicode.ToUpper(new[i])
			b=true
		}else if unicode.IsLetter(c) && b{
			new[i]=unicode.ToLower(new[i])
		}
	}
	return string(new)
}