package sidefunctions



func Punc(str []string) []string {

	for i := 0; i < len(str); i++ {
		if i != 0 && !NotPunc(str[i]) {
			str[i-1] = str[i-1] + str[i]
					str[i] = ""
					str = Clean(SliceToString(str))
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
