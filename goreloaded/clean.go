package goreloaded

import (
	"strings"
)

func Clean(s string) []string {
	str := ""
	var cleaned []string
	paret := false
	for i, c := range s {
		if c == '(' {
			for j := i; j < len(s); j++ {
				if s[j] == ')' {
					if IsFlag("("+strings.Join((Clean(s[i+1:j])), "")+")") || IsMultiFlag("("+strings.Join((Clean(s[i+1:j])), "")+")") {
						if !paret && str != "" {

							cleaned = append(cleaned, str)
							str = ""
						}
						paret = true
						break
					}
				}
			}
		}

		if paret && c == ')' {

			cleaned = append(cleaned, "("+strings.TrimSpace(str[1:])+")")
			str = ""
			paret = false
			continue
		}
		if i != len(s)-1 && c == '\n' {

			cleaned = append(cleaned, str+"\n")

			continue
		}
		if c != ' ' && i != len(s)-1 {
			str += string(c)
			continue
		}
		if str != "" && c == ' ' && !paret {

			cleaned = append(cleaned, str)
			str = ""
		}
		if i == len(s)-1 && c != ' ' && !paret {
			cleaned = append(cleaned, str+string(c))
		}
	}

	return cleaned
}
