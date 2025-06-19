package sidefunctions

func Clean(s string) []string {
	cou := 0
	var cleaned []string
	b := false
	for i, c := range s {
		if c == '(' {
			b = true
		}
		if c == ')' {
			b = false
		}
		if c==',' && cou!=0 &&!b {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i])
			cou = 0
		}
		if c != ' ' && i != len(s)-1 && !b {
			cou++
			continue
		}else if b {
			cou++
		}
		if cou != 0 && c == ' ' && !b {
			cou = i - cou
			cleaned = append(cleaned, s[cou:i])
			cou = 0
		}
		if i == len(s)-1 && c != ' ' && !b {
			cou = i - cou
			cleaned = append(cleaned, s[cou:])
		}
	}
	return cleaned
}
