package goreloaded

import "strings"

func Punc(str []string) []string {
	for i := 0; i < len(str); i++ {
		if Paret(str[i]) {
			str[i] = "(" +strings.Join(Punc(Clean(str[i][1:len(str[i])-1]))," ") + ")"
		}
		if i != 0 && AllPunc(str[i]) {
			str[i-1] =strings.Join( Clean(str[i-1])," ") + str[i]
			str[i] = ""
			str = Clean(strings.Join(str, " "))
			i--

		}

	}
	return str
}

func IsPunc(s rune) bool {
	if s == ',' || s == '.' || s == ':' || s == ';' || s == '?' || s == '!' {
		return true
	}
	return false
}

func AllPunc(s string) bool {
	cou := 0
	for _, c := range s {
		if c == ',' || c == '.' || c == ':' || c == ';' || c == '?' || c == '!' {
			cou++
		}
	}
	if cou == len(s) {
		return true
	} else {
		return false
	}
}
