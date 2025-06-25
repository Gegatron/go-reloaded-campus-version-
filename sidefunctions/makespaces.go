package sidefunctions

import "unicode"

func MakeSpaces(s string)string{
b:=false

new:=""
	for i,c:=range s{
		if c=='(' {
			for j := i; j < len(s); j++ {
				if s[j]==')' {
					if IsFlag("("+SliceToString(Punc(Clean(s[i+1:j])))+")") || IsMultiFlag("("+SliceToString(Punc(Clean(s[i+1:j])))+")"){
						new=new+" "+string(c)
						b=true
						
						break
					}
				}
			}
			
			
			if !b {
				new=new+string(c)
			}
			
			
			
			

			continue
		}
		 if c==')' && b {
			
			new=new+string(c)+" "
			b=false
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