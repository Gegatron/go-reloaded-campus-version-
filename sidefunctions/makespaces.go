package sidefunctions

func MakeSpaces(s string)string{

b:=false


new:=""
	for i := 0; i < len(s); i++ {
		if !b &&s[i]=='(' {
			
			new=new+" "+string(s[i])
			
			b=true
			
			

			continue
		}
		 if b && s[i]==')' {
			new=new+string(s[i])+" "
			b=false
			continue
		}else if !b && IsPunc(s[i]){
			new=new+" "+string(s[i])+" "
		}else{
			new=new+string(s[i])
		}
		
	}
	return new
}