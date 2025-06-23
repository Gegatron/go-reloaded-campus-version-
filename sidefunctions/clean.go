package sidefunctions

func Clean(s string) []string {
	cou := 0
	var cleaned []string
	b := 0

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			b++
		}
		if s[i] == ')' {
			b--
		}

		if s[i] != ' ' && i != len(s)-1 && b==0 {
			cou++
			continue
		} else if b>0 {
			cou++
		}
		if cou != 0 && s[i] == ' ' && b==0 {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i])
			cou = 0
		}
		if i == len(s)-1 && s[i] != ' ' && b==0 {
			cou = i - cou
			cleaned = append(cleaned, s[cou:])
		}
	}
	return cleaned
}
