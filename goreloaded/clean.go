package goreloaded

import (
	
	"strings"
)

func Clean(s string) []string {
	cou := 0
	var cleaned []string
	paret := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			for j := i; j < len(s); j++ {
				if s[j] == ')' {
					if IsFlag("("+strings.Join(Punc(FixFlags(Clean(s[i+1:j])))," ")+")") || IsMultiFlag("("+strings.Join(Punc(FixFlags(Clean(s[i+1:j])))," ")+")") {
						paret++
						break
					} else {
						continue
					}
				}
			}
		}
		if paret > 0 && s[i] == ')' {
			paret--
		}
		if i!=len(s)-3 && s[i] == '\n'  {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i]+"\n")
			cou = 0
			continue
		}
		if s[i] != ' ' && i != len(s)-1 && paret== 0 {
			cou++
			continue
		} else if paret > 0 {
			cou++
		}
		if cou != 0 && s[i] == ' ' && paret == 0 {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i])
			cou = 0
		}
		if i == len(s)-1 && s[i] != ' ' && paret == 0 {
			cou = i - cou
			cleaned = append(cleaned, s[cou:])
		}
	}
	
	return cleaned
}
