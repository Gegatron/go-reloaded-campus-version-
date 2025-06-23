package sidefunctions

import (
	"strings"
	"unicode"
)

func Reload(str string) []string {
	c := Clean(str)
	n := 0
	s := ""
	for i := 0; i < len(c); i++ {
		if strings.Contains(c[i], "(up,") {
			n , err := GetNumber(c[i])
			if err!=nil {
				c[i]="(up, <Invalid Number>)"
					continue
			}
			for j := i; j >= 0; j-- {
				if IsFlag(c[j]) {
					continue
				}
				if n >= 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							c[j] = strings.ToUpper(c[j])
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		} else if strings.Contains(c[i], "(low,") {
			n , err := GetNumber(c[i])
			if err!=nil {
				c[i]="(low, <Invalid Number>)"
				continue
			}
			for j := i; j >= 0; j-- {
				if IsFlag(c[j]) {
					continue
				}
				if n >= 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							c[j] = strings.ToLower(c[j])
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		} else if strings.Contains(c[i], "(cap,") {
			n , err := GetNumber(c[i])
			if err!=nil {
				c[i]="(cap, <Invalid Number>)"
					continue
			}
			for j := i; j >= 0; j-- {
					if IsFlag(c[j]) {
					continue
				}
				if n >= 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							c[j] = Capitalize(c[j])
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		}
		switch c[i] {
		case "(up)":
			n = 1
			if n > i {
				n = i
			}
			for j := i; j >= 0; j-- {
				if IsFlag(c[j]) {
					continue
				}
				if n >= 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							c[j] = strings.ToUpper(c[j])
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		case "(low)":
			n = 1
			if n > i {
				n = i
			}
			for j := i; j >= 0; j-- {
				if IsFlag(c[j]) {
					continue
				}
				if  n >= 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							c[j] = strings.ToLower(c[j])
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		case "(cap)":
			n = 1
			if n > i {
				n = i
			}
			for j := i; j >= 0; j-- {
				if IsFlag(c[j]) {
					continue
				}
				if  n >= 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							c[j] = Capitalize(c[j])
					n--
					break
						}
					}
					
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue

		case "(hex)":

			c[i-1] = ToHex(c[i-1])
			c[i] = ""
			s = SliceToString(c)
			c = Clean(s)
			i--

		case "(bin)":

			c[i-1] = ToBin(c[i-1])
			c[i] = ""
			s = SliceToString(c)
			c = Clean(s)
			i--
		default:

		}

	}


	return c
}
