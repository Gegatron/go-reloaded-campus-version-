package sidefunctions

func ToUpper(s string) string { 
	new:=[]rune(s)
	for i, c:=range s 	{
		if c>='a' && c<='z' {
			new[i]=c-32
		}
		
	}
	return string(new)
}
