package goreloaded





func Trait(s string) string {
	proto := MakeSpaces(string(s))
	fixed := Clean(proto)
	fixed = FixFlags(fixed)
	
	fixed = Punc(fixed)
	
	fixed = Quotes(fixed)
	
	fixed = ATooAn(fixed)
	str := ""
	for i := 0; i < len(fixed); i++ {
		if fixed[i][len(fixed[i])-1] == '\n' || i==len(fixed)-1 {
			str += fixed[i] 
			continue
		}
		str += fixed[i] + " "
	}

	return str
}
