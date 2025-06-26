package sidefunctions

func Cleannewlines(s string)string{
	proto:=[]string{}
	cou:=0
	for i,c:=range s {
		if c!='\n'{
			cou++
			continue
		}
		if c=='\n'  {
			
			cou=i-cou
			proto=append(proto, s[cou:i])
			cou=0
			continue
		} 
	}
	cleaned:=""
	for j := 0; j < len(proto); j++ {
		if proto[j]!="" {
			cleaned+=proto[j]+"\n"
		}
		
	}
	return cleaned
}