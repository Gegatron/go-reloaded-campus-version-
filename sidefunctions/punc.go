package sidefunctions

func Punc(str []string)[]string{
	cou:=0
	for i := 0; i < len(str); i++ {
		
		for j := 0; j < len(str[i]); j++ {
			if IsPunc(str[i][j]) {
				cou++
			}else{
				break
			}
		}
		if i==0 {
			continue
		}
		str[i-1]=str[i-1]+str[i][:cou]
		str[i]=str[i][cou:]
		cou=0
	}
	return str
}
func IsPunc(s byte)bool{
	if s==',' || s=='.'|| s==':'|| s==';'|| s=='?'|| s=='!' {
		return true

	}
	return false
}