package goreloaded

import (

	"strconv"
	"strings"
	"unicode"
)

func FixFlags(c []string) []string {
	
	n := 0
	c=Clean(strings.Join(c," "))
	for i := 0; i < len(c); i++ {
	
			
			if IsFlag(c[i]) {
			n = 1
			if n > i {
				n = i
			}
			for j := i - 1; j >= 0; j-- {
				if c[j][len(c[j])-1]=='\n' {
					break
				}
				if n > 0 {
					for _,u:=range c[j] {
						 if unicode.IsLetter(u) || unicode.IsNumber(u) {
							if c[i] == "(up)" {
								if unicode.IsNumber(u) {
									continue
								}
								c[j] = strings.ToUpper(c[j])
							} else if c[i] == "(low)" {
								if unicode.IsNumber(u ){
									continue
								}
								c[j] = strings.ToLower(c[j])
							} else if c[i] == "(cap)" {
								if unicode.IsNumber(u) {
									continue
								}
								c[j] = Capitalize(c[j])
							}else if c[i]=="(bin)" {
								c[j] = ToBin(c[j])
							}else if c[i]=="(hex)" {
								c[j] = ToHex(c[j])
								
							}
						}else{
							continue
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
			i=0
			continue
		}
			if IsMultiFlag(strings.Join(Punc(Clean(c[i])), " ")) {
				n, err := GetNumber(c[i])
				if err != nil {
					continue
				}
				if n<0 {
					n=0
				}
				for j := i - 1; j >= 0; j-- {
					if c[j][len(c[j])-1]=='\n' {
					break
				}
					if n > 0 {
						for _,u:=range c[j] {
							if unicode.IsLetter(u) {
								if strings.HasPrefix(strings.Join(Punc(Clean(c[i])), " "), "(up,") {
									c[j] = strings.ToUpper(c[j])
								} else if strings.HasPrefix(strings.Join(Punc(Clean(c[i])), " "), "(low,") {
									c[j] = strings.ToLower(c[j])
								} else if strings.HasPrefix(strings.Join(Punc(Clean(c[i])), " "), "(cap,"){
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
				i=0
				continue
			}
		
	}
	return c
}

func GetNumber(s string) (int, error) {
	new := ""
	b := false
	for i := 0; i < len(s); i++ {
		if !b && s[i] == '-' && new == "" {
			new+=string(s[i])
			b=true
			continue
		}
		if !b && (s[i] >='0'  && s[i]<='9') && new==""{
			new+=string(s[i])
			b=true
			continue
		}
		if b && i != len(s)-1 {
			new +=string(s[i])
		}
		
	}
	
	return strconv.Atoi(new)
}

func Capitalize(s string)string{
	new:=[]rune(s)
	b:=false
	for i,c:=range new{
		if IsLetter(c) && !b{
			new[i]=unicode.ToUpper(new[i])
			b=true
		}else if IsLetter(c) && b{
			new[i]=unicode.ToLower(new[i])
		}
	}
	return string(new)
}

func IsLetter(s rune) bool {
	return unicode.Is(unicode.Latin, (s)) && unicode.IsLetter(s)
}

func ToHex(s string)string{
	b:=false
	if s[len(s)-1]=='\n' {
		b=true
		s=s[:len(s)-1]
	}
	hex,err:= strconv.ParseInt(s,16,64)
if err!=nil {
	return s
}
if b {
	return strconv.Itoa(int(hex))+"\n"
}
return strconv.Itoa(int(hex))
}

func ToBin(s string)string{
	b:=false
	if s[len(s)-1]=='\n' {
		b=true
		s=s[:len(s)-1]
	}
bin,err:= strconv.ParseInt(s,2,64)
if err!=nil {
	return s
}
if b {
	return strconv.Itoa(int(bin))+"\n"
}
return strconv.Itoa(int(bin))
}

func Paret(s string)bool{
if s[0]=='(' && s[len(s)-1]==')' {
	return true
}
return false
}