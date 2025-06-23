package sidefunctions



func IsMultiFlag(s string)bool{
	if s[0]=='(' && s[len(s)-1]==')' {
		if BeginsWith(s,"(up,")||BeginsWith(s,"(cap,")||BeginsWith(s,"(low,") {
			return true
		}
	}
	return false
}
func IsFlag(s string)bool {

		if s=="(up)" || s=="(low)" || s=="(cap)"  {
			return true
		}
	
	return false
}