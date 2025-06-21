package sidefunctions

func NotQuotes(s string) bool {
	cou:=0
	for _,c:=range s{
		if c=='\'' {
			cou++
		}
	}
	if cou==len(s) {
		return false
	}else{
		return true
	}
}
