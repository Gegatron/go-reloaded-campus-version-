package sidefunctions

func GetNumber(s string) int {
	num := 0
	for _, c := range s {
		if c == '-' {
			return 0
		}
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
	}
	return num
}