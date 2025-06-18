package sidefunctions

func clean(s string)[]string{
	cou:=0
	var cleaned []string
	for i,c:=range s{
		if c!=' ' && i!=len(s)-1{
			cou++
			continue
		}
		if  cou!=0 && c==' '{
			cou=i-cou
			cleaned = append(cleaned, s[cou:i])
			cou=0
		}
		if i==len(s)-1 && c!=' ' {
			cou=i-cou
			cleaned = append(cleaned, s[cou:])
		}
	}
	return cleaned
}