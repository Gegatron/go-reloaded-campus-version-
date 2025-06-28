package goreloaded

import (
	
	"strings"
)

func Quotes(reloaded []string) []string {
	s := strings.Join(reloaded, " ")
	index := 0
	str := ""
	b := true
	for i ,c:=range s {
		if !b &&c=='\n' {
			str += "'" + strings.Join(Clean(s[index+1:i]), " ")
			b=true
		}
		if b &&c == '\'' {
			if i == len(s)-1 {
				str += string(c)
				continue
			}
			if i==0 {
				str += " "
			index = i
			b = false
				continue
			}
			if s[i+1]==' ' || s[i-1]==' ' {
				str += " "
			index = i
			b = false
			continue
			}
			
		}
		if !b && c == '\'' {
			str += "'" + strings.Join(Clean(s[index+1:i]), " ") + "'" + " "
			b = true
			continue
		}
		if b {
			str += string(c)
		}
		if !b && i == len(s)-1 {
			str += s[index:]
		}
	}
	
	reloaded = Clean(str)
	return reloaded
}
