package sidefunctions

func ToLower(s string) string { 
	new:=[]rune(s)
	for i, c:=range s 	{
		if c>='A' && c<='Z' {
			new[i]=c+32
		}
		
	}
	return string(new)
}
