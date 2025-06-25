package sidefunctions



func IsMultiFlag(s string)bool{
	b:=false
	if s[0]=='(' && s[len(s)-1]==')' {
		
		if BeginsWith(s,"(up,") {
			for i := 4; i < len(s)-1; i++ {
				if s[i]!=' ' {
					b=true
				}
				if b&&(s[i]<'0'||s[i]>'9'){
					return false
				}
			}
			return true
		}
		if BeginsWith(s,"(low,") {
			for i := 5; i < len(s)-1; i++ {
				if s[i]!=' ' {
					b=true
				}
				if b&&(s[i]<'0'||s[i]>'9'){
					return false
				}
			}
			return true
		}
		if BeginsWith(s,"(cap,") {
			for i := 5; i < len(s)-1; i++ {
				if s[i]!=' ' {
					b=true
				}
				if b&&(s[i]<'0'||s[i]>'9'){
					return false
				}
			}
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