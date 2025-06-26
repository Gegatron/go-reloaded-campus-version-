package sidefunctions

import (
	
	"strings"
)



func Trait(s string)string{
proto := MakeSpaces(string(s))

	
	fixed:=Clean(proto)
	
	fixed=Reload(fixed)
	
	fixed = Punc(fixed)
	
	fixed = Quotes(fixed)
	
	

	fixed = ATooAn(fixed)
	/*fmt.Println(fixed)
	str:=""
	for i := 0; i < len(fixed); i++ {
		if fixed[i]=="\n" {
			continue
		}
		if i!=len(fixed)-1  && fixed[i+1]!="" {
			str+=fixed[i]
			continue
		}else if i==len(fixed)-1{
			str+=fixed[i]
			continue
		}
		str+=fixed[i]+" "
	}*/
	
	return strings.Join(fixed," ")
}

