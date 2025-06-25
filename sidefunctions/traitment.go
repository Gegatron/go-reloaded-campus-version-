package sidefunctions

func Trait(s string)string{
proto := MakeSpaces(string(s))
	

	
	fixed:=Clean(proto)
	
	fixed=Reload(fixed)
	
	fixed = Punc(fixed)
	
	fixed = Quotes(fixed)
	
	

	fixed = ATooAn(fixed)
	return SliceToString(fixed)
}

