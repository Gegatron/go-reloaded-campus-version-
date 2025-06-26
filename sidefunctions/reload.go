package sidefunctions

import (

	"strconv"
	"strings"
	"unicode"
)

func Reload(c []string) []string {
	temp := ""
	n := 0
	c=Clean(strings.Join(c," "))
	for i := 0; i < len(c); i++ {
			
		if IsFlag(c[i]) {
			n = 1
			if n > i {
				n = i
			}
			for j := i - 1; j >= 0; j-- {
				if n > 0 {
					for u := 0; u < len(c[j]); u++ {
						if unicode.IsLetter(rune(c[j][u])) {
							if c[i] == "(up)" {
								c[j] = strings.ToUpper(c[j])
							} else if c[i] == "(low)" {
								c[j] = strings.ToLower(c[j])
							} else if c[i] == "(cap)" {
								c[j] = Capitalize(c[j])
							}

							
						}else  {
							if c[i]=="(bin)" {
								c[j] = ToBin(c[j])
							}else if c[i]=="(hex)" {
								c[j] = ToHex(c[j])
							}
						}
						n--
							break
					}
				} else {
					break
				}
			}
			c[i] = ""
			c = Clean(strings.Join(c, " "))
			i--
			continue
		}
	

		if Paret(c[i]) {

			if !IsMultiFlag(strings.Join(Punc(Clean(c[i])), " ")) {
				if c[i] != temp {
					temp = c[i]
					c[i] = "(" + strings.Join(Reload(Clean(c[i][1:len(c[i])-1])), " ") + ")"

				}
			}

			if IsMultiFlag(strings.Join(Punc(Clean(c[i])), " ")) {
				n, err := GetNumber(c[i])
				if err != nil {
					c[i] = ""
					c = Clean(strings.Join(c, " "))
					continue
				}
				for j := i - 1; j >= 0; j-- {
					if n > 0 {
						for u := 0; u < len(c[j]); u++ {
							if unicode.IsLetter(rune(c[j][u])) {
								if strings.HasPrefix(strings.Join(Punc(Clean(c[i])), " "), "(up,") {
									c[j] = strings.ToUpper(c[j])
								} else if strings.HasPrefix(strings.Join(Punc(Clean(c[i])), " "), "(low,") {
									c[j] = strings.ToLower(c[j])
								} else if strings.HasPrefix(strings.Join(Punc(Clean(c[i])), " "), MakeSpaces("(cap,")) {
									c[j] = Capitalize(c[j])
								}
								n--
								break
							}
						}
					} else {
						break
					}
				}
				c[i] = ""
				c = Clean(strings.Join(c, " "))
				i--
				continue
			}

		}
	}

	return c
}

func GetNumber(s string) (int, error) {
	new := ""
	b := false
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && new == "" {
			return strconv.Atoi("0")
		}
		if s[i] >= '0' && s[i] <= '9' {
			b = true
		}
		if b && !(s[i] >= '0' && s[i] <= '9') && i != len(s)-1 {
			new = ""
			break
		}
		if b && i != len(s)-1 {
			new = new + string(s[i])
		}
	}
	
	return strconv.Atoi(new)
}
