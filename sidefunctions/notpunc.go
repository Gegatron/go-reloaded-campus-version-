package sidefunctions

func NotPunc(s string)bool{
	cou:=0
	for _,c:=range s{
		if c==',' || c=='.' || c==':' || c==';' || c=='?' || c=='!'  {
			cou++
		}
	}
	if cou==len(s) {
		return false
	}else{
		return true
	}
}