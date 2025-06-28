package goreloaded

import "strings"

func MakeSpaces(s string) string {
	b := 0

	str := ""
	for i, c := range s {
		if c == '(' {
			for j := i; j < len(s); j++ {
				if s[j] == ')' {
					if IsFlag("("+strings.Join(Punc(FixFlags(Clean(MakeSpaces(s[i+1:j]))))," ")+")") || IsMultiFlag("("+strings.Join(Punc(FixFlags(Clean(MakeSpaces(s[i+1:j]))))," ")+")") {
						str = str + " " + string(c)
						b++

						break
					}
				}
			}

			if b == 0 {
				str = str + string(c)
			}

			continue
		}
		if c == ')' && b > 0 {

			str = str + string(c) + " "
			b--
			continue
		} else if i != len(s)-1 && i != 0 && IsPunc(c) && s[i+1] != ' ' && s[i-1] == ' ' {
			str = str + string(c) + " "
		} else if i != len(s)-1 && i != 0 && IsPunc(c) && (IsPunc(rune(s[i-1]))) {
			str = str + string(c) + " "
		} else {
			str = str + string(c)
		}

	}
	return str
}
