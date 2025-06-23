package sidefunctions

import "unicode"

func MakeSpaces(s string)string{

b:=false


new:=""
	for i,c:=range s{
		if !b &&c=='(' {
			
			new=new+" "+string(c)
			
			b=true
			
			

			continue
		}
		 if b && c==')' {
			new=new+string(c)+" "
			b=false
			continue
		}else if !b && IsPunc(c){
			new=new+" "+string(c)+" "
		}else if !b && i!=0 && i!=len(s)-1 && IsQuote(c) && !(unicode.IsLetter(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1]))){
			new=new+" "+string(c)+" "
		}else if !b && i==0 && IsQuote(c){
			new=new+string(c)+" "
		}else if  !b && i==len(s)-1 && IsQuote(c){
			new=new+" "+string(c)
		}else{
			new=new+string(c)
		}
		
		
		
	}
	return new
}