package sidefunctions

import "strings"

func IsFlag(s string)bool{
	if s[0]=='(' && s[len(s)-1]==')' {
		if strings.Contains(s,"up")||strings.Contains(s,"cap")||strings.Contains(s,"low")||strings.Contains(s,"hex")||strings.Contains(s,"bin") {
			return true
		}
	}
	return false
}