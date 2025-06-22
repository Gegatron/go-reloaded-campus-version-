package sidefunctions

import "unicode"

func ATooAn(s []string) []string {
	b:=false
	var n int
	for i := 0; i < len(s); i++ {
		for u := 0; u < len(s[i]); u++ {
			if (s[i][u] == 'a' || s[i][u] == 'A'){
				if u==len(s[i])-1 {
					n=u
					b=true
					break
				}else{
					break
				}
			}else if  !unicode.IsLetter(rune(s[i][u])) {
				continue
			}else{
				break
			}	
			
			
		
		}
		if b {
				for j := 0; j < len(s[i+1]); j++ {
				if unicode.IsLetter(rune(s[i+1][j])) && IsVowel(rune(s[i+1][j])) {
					if s[i][n]=='a' {
						s[i]=s[i][:n]+"an"
						break
					}else if s[i][n]=='A' {
						s[i]=s[i][:n]+"An"
						break
					}
					
				}else if unicode.IsLetter(rune(s[i+1][j])) && !IsVowel(rune(s[i+1][j])) {
						break
					}
			}
			b=false
			}
	}
	return s
}
func IsVowel(c rune)bool{
	if c=='a'||c=='e'||c=='o'||c=='i'||c=='u'||c=='h'||c=='A'||c=='E'||c=='O'||c=='I'||c=='U'||c=='H' {
		return true
	}
	return false
}