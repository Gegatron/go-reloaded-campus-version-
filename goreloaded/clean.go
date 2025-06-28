package goreloaded

import (
	
	"strings"
)

func Clean(s string) []string {
	cou := 0
	var cleaned []string
	paret := false
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			for j := i; j < len(s); j++ {
				if s[j] == ')' {
					if IsFlag("("+strings.Join(Punc((Clean(s[i+1:j])))," ")+")") || IsMultiFlag("("+strings.Join(Punc((Clean(s[i+1:j])))," ")+")") {
						paret=true
						break
					} 
				}
			}
		}
		if paret  && s[i] == ')' {
			paret=false
		}
		if i!=len(s)-1 && s[i] == '\n'  {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i]+"\n")
			cou = 0
			continue
		}
		if s[i] != ' ' && i != len(s)-1 && !paret {
			cou++
			continue
		} else if paret {
			cou++
		}
		if cou != 0 && s[i] == ' ' && !paret {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i])
			cou = 0
		}
		if i == len(s)-1 && s[i] != ' ' && !paret{
			cou = i - cou
			cleaned = append(cleaned, s[cou:])
		}
	}
	
	return cleaned
}
