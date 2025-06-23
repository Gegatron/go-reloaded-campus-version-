package sidefunctions

func Clean(s string) []string {
	cou := 0
	var cleaned []string
	b := false

	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			b = true
		}
		if s[i] == ')' {
			b = false
		}

		if s[i] != ' ' && i != len(s)-1 && !b {
			cou++
			continue
		} else if b {
			cou++
		}
		if cou != 0 && s[i] == ' ' && !b {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i])
			cou = 0
		}
		if i == len(s)-1 && s[i] != ' ' && !b {
			cou = i - cou
			cleaned = append(cleaned, s[cou:])
		}
	}
	return cleaned
}
