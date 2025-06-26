package sidefunctions

import (
	"strings"
	"unicode"
)

func Reload(c []string) []string {
	temp:=""
	n := 0
	s := ""
	for i := 0; i < len(c); i++ {
		if IsFlag(c[i]) {
			n = 1
			if n > i {
				n = i
			}
			for j := i-1; j >= 0; j-- {
				
				if n > 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							if c[i]=="(up)" {
								c[j] = strings.ToUpper(c[j])
							}else if c[i]=="(low)" {
								c[j] = strings.ToLower(c[j])
							}else if c[i]=="(cap)" {
								c[j] = Capitalize(c[j])
							}
							
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		}
		switch c[i] {
	
		case "(hex)":

			c[i-1] = ToHex(c[i-1])
			c[i] = ""
			s = SliceToString(c)
			c = Clean(s)
			i--

		case "(bin)":

			c[i-1] = ToBin(c[i-1])
			c[i] = ""
			s = SliceToString(c)
			c = Clean(s)
			i--
		default:

		}

		if Paret(c[i]) {
		
				if !IsMultiFlag(SliceToString(Punc(Clean(c[i]))) ) {
				
				if c[i]!=temp {
					temp=c[i]
					c[i]="("+SliceToString(Reload(Clean(c[i][1:len(c[i])-1])))+")"
				
				
				}
				
				
			}
		
		 if IsMultiFlag(SliceToString(Punc(Clean(c[i]))) ){
				n , err := GetNumber(c[i])
				
			if err!=nil {
				c[i]=""
				c=Clean(SliceToString(c))
					continue
			}
			for j := i-1; j >= 0; j-- {
				
				if n > 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							if BeginsWith(SliceToString(Punc(Clean(c[i]))),"(up,") {
								c[j] = strings.ToUpper(c[j])
							}else if BeginsWith(SliceToString(Punc(Clean(c[i]))),"(low,") {
								c[j] = strings.ToLower(c[j])
							}else if BeginsWith(SliceToString(Punc(Clean(c[i]))),MakeSpaces("(cap,")){
								c[j] = Capitalize(c[j])
							}
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
			}
			
		
		

		}
	}


	return c
}
