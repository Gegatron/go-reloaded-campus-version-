package sidefunctions

import "unicode"

func MakeSpaces(s string)string{


new:=""
	for i,c:=range s{
		if c=='(' {
			
			new=new+" "+string(c)
			
			
			
			

			continue
		}
		 if c==')' {
			new=new+string(c)+" "
			
			continue
		}else if  IsPunc(c){
			new=new+" "+string(c)+" "
		}else if  i!=0 && i!=len(s)-1 && IsQuote(c) && !(unicode.IsLetter(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1]))){
			new=new+" "+string(c)+" "
		}else if  i==0 && IsQuote(c){
			new=new+string(c)+" "
		}else if   i==len(s)-1 && IsQuote(c){
			new=new+" "+string(c)
		}else{
			new=new+string(c)
		}
		
		
		
	}
	return new
}