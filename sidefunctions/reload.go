package sidefunctions

import (
	"strings"
)

func Reload(str string) []string {
	c := Clean(str)
	n := 0
	s := ""
	for i := 0; i < len(c); i++ {
		if strings.Contains(c[i], "(up,") {
			n = GetNumber(c[i])
			for j := i; j >= 0; j-- {
				if n >= 0 {
					c[j] = strings.ToUpper(c[j])
					n--
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		} else if strings.Contains(c[i], "(low,") {
			n = GetNumber(c[i])
			for j := i; j >= 0; j-- {
				if n >= 0 {
					c[j] = strings.ToLower(c[j])
					n--
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(SliceToString(c))
			i--
			continue
		} else if strings.Contains(c[i], "(cap,") {
			n = GetNumber(c[i])
			for j := i; j >= 0; j-- {
				if n >= 0 {
					c[j] = Capitalize(c[j])
					n--
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
				if n >= 0 {
					c[j] = strings.ToUpper(c[j])
					n--
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
				if NotPunc(c[j]) && NotQuotes(c[j]) && n >= 0 {
					c[j] = strings.ToLower(c[j])
					n--
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
				if NotPunc(c[j]) && NotQuotes(c[j]) && n >= 0 {
					c[j] = Capitalize(c[j])
					n--
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

	c=Quotes(c)


	return c
}
