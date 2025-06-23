package sidefunctions

func Paret(s string)bool{
if s[0]=='(' && s[len(s)-1]==')' {
	return true
}
return false
}