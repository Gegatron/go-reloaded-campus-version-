package goreloaded

import (
	"strings"
)

func ATooAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == "a" && s[i+1][len(s[i+1])-1] != '\'' {
			if strings.Contains("aeoiuhAEOIUH", string(s[i+1][0])) {
				s[i] = "an"
			}
		}
		if i != len(s)-1 && s[i] == "A" && s[i+1][len(s[i+1])-1] != '\'' {
			if strings.Contains("aeoiuhAEOIUH", string(s[i+1][0])) {
				s[i] = "An"
			}
		}
		if i != len(s)-1 && s[i] == "'a" && s[i+1][len(s[i+1])-1] == '\'' {
			if strings.Contains("aeoiuhAEOIUH", string(s[i+1][0])) {
				s[i] = "'an"
			}
		}
		if i != len(s)-1 && s[i] == "'A" && s[i+1][len(s[i+1])-1] == '\'' {
			if strings.Contains("aeoiuhAEOIUH", string(s[i+1][0])) {
				s[i] = "'An"
			}
		}
	}
	s = Clean(strings.Join(s, " "))
	return s
}
