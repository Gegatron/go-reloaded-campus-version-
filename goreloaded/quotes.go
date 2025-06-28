package goreloaded

import "strings"

func Quotes(reloaded []string) []string {
	s := strings.Join(reloaded, " ")
	index := 0
	str := ""
	b := true
	for i := 0; i < len(s); i++ {
		if b && s[i] == '\'' {
			if i == len(s)-1 {
				str += string(s[i])
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
		if !b && s[i] == '\'' {
			str += "'" + strings.Join(Clean(s[index+1:i]), " ") + "'" + " "
			b = true
			continue
		}
		if b {
			str += string(s[i])
		}
		if !b && i == len(s)-1 {
			str += s[index:]
		}
	}
	reloaded = Clean(str)
	return reloaded
}
