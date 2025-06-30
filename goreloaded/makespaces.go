package goreloaded

func MakeSpaces(s string) string {
	str := ""
	for i, c := range s {
		if i != len(s)-1 && IsPunc(c) && s[i+1] != ' ' {
			str = str + string(c) + " "
		} else {
			str = str + string(c)
		}
	}
	return str
}
