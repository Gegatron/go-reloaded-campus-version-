package sidefunctions

import (
	"strings"
	"unicode"
)

func Reload(str string) []string {
	c := Clean(str)
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
		if IsMultiFlag(c[i]) {
			c[i]="("+SliceToString(FixFlags(c[i]))+")"
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
							if BeginsWith(c[i],"(up,") {
								c[j] = strings.ToUpper(c[j])
							}else if BeginsWith(c[i],"(low,") {
								c[j] = strings.ToLower(c[j])
							}else if BeginsWith(c[i],"(cap,"){
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

	}


	return c
}
