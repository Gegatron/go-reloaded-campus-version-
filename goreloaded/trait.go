package goreloaded


func Trait(s string) string {
	proto := MakeSpaces(string(s))
	fixed := Clean(proto)
	
	fixed = FixFlags(fixed)
	
	fixed = Punc(fixed)
	
	fixed = Quotes(fixed)
	
	fixed = Punc(fixed)
	
	fixed = ATooAn(fixed)
	
	str := ""
	for i := 0; i < len(fixed); i++ {
		if i == 0 && fixed[i] == "\n" {
			continue
		}
		if fixed[i] == "\n" && fixed[i-1][len(fixed[i-1])-1] == '\n' {
			continue
		}

		if i != len(fixed)-1 && fixed[i+1] != "" && fixed[i][len(fixed[i])-1] == '\n' {
			str += fixed[i]
			continue
		} else if i == len(fixed)-1 {
			str += fixed[i]
			continue
		}
		str += fixed[i] + " "
	}
	

	return str
}
