package sidefunctions

func SliceToString(sli []string)string{
	str:=""
	for i,v:=range sli{
		if i!=len(sli)-1 {
			str=str+v+" "
			continue
		}
		str=str+v
	}
	return str
}