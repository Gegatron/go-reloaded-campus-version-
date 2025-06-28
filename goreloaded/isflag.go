package goreloaded

import "strings"

func IsMultiFlag(s string) bool {
	b := false
	if s[0] == '(' && s[len(s)-1] == ')' {

		if strings.HasPrefix(s, "(up,") {
			for i := 4; i < len(s)-1; i++ {
				if s[i] != ' ' && !b {
					if s[i] == '-' || s[i] == '+' {
						b = true
						continue
					}
					b = true
				}
				if b && (s[i] < '0' || s[i] > '9') {
					return false
				}
			}
			return true
		}
		if strings.HasPrefix(s, "(low,") {
			for i := 5; i < len(s)-1; i++ {
				if s[i] != ' ' && !b {
					if s[i] == '-' || s[i] == '+' {
						b = true
						continue
					}
					b = true
				}
				if b && (s[i] < '0' || s[i] > '9') {
					return false
				}
			}
			return true
		}
		if strings.HasPrefix(s, "(cap,") {
			for i := 5; i < len(s)-1; i++ {
				if s[i] != ' ' && !b {
					if s[i] == '-' || s[i] == '+' {
						b = true
						continue
					}
					b = true
				}
				if b && (s[i] < '0' || s[i] > '9') {
					return false
				}
			}
			return true
		}
	}
	return false
}

func IsFlag(s string) bool {
	if s == "(up)" || s == "(low)" || s == "(cap)" || s=="(bin)" || s=="(hex)" {
		return true
	}

	return false
}
