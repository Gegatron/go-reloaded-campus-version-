package sidefunctions



func MakeSpaces(s string)string{
b:=0

new:=""
	for i,c:=range s{
		if c=='(' {
			for j := i; j < len(s); j++ {
				if s[j]==')' {
					if IsFlag("("+SliceToString(Punc(Reload(Clean(s[i+1:j]))))+")") || IsMultiFlag("("+SliceToString(Punc(Reload(Clean(s[i+1:j]))))+")") {
						new=new+" "+string(c)
						b++
						
						break
					}
				}
			}
			
			
			if b==0 {
				new=new+string(c)
			}
			
			
			
			

			continue
		}
		 if c==')' && b>0 {
			
			new=new+string(c)+" "
			b--
			continue
		}else if i!=len(s)-1 && i!=0 && IsPunc(c) && s[i+1]!=' ' && s[i-1]==' ' {
			new=new+string(c)+" "
		}else{
			new=new+string(c)
		}
		
		
		
	}
	return new
}