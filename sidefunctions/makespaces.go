package sidefunctions

func MakeSpaces(s string)string{

b:=false
n:=0
b2:=false
	for i := 0; i < len(s); i++ {
		if s[i]=='(' {
			if !b2 {
				n=i
				b=true
			}
			s=s[:n]+" "+s[n:]
			b=true
			n++
			

			continue
		}
		 if b && s[i]==')' {
			s=s[:n]+" "+s[n:]
			b=false
			n++
			continue
		}
		
	}
	return s
}