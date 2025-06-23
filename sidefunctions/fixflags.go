package sidefunctions

import (
	"strings"
	
)

func FixFlags(s string)[]string{

	new:=Clean(s[1:len(s)-2])
	for i := 0; i < len(new); i++ {
		if IsFlag(new[i]) {
		
				
				
				
							if new[i]=="(up)" {
								new[i-1] = strings.ToUpper(new[i-1])
							}else if new[i]=="(low)" {
								new[i-1] = strings.ToLower(new[i-1])
							}else if new[i]=="(cap)" {
								new[i-1] = Capitalize(new[i-1])
							}else{
								new[i-1]=""
							}
							
					
					
					
					
			
			new[i]=""
			new=Clean(SliceToString(new))
			i--
			continue
		}
		if IsMultiFlag(new[i]) {
			new[i]=SliceToString(FixFlags(new[i]))
			
			new=Clean(SliceToString(new))
			i--
		}
	}
	return new
}