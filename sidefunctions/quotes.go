package sidefunctions

import "strings"

func Quotes(reloaded []string)[]string{
	var quotes []string
	str:=SliceToString(reloaded)
	cou:=0
	b:=0
	for i,c:=range str {
		if i==len(str)-1 && c!='\'' {
			if cou==0 {
				return reloaded
			}
			
			quotes = append(quotes, "'"+str[cou:])
			cou=i+1
		}
		if i==len(str)-1 && c=='\'' && b%2==0{
			quotes = append(quotes, str[cou:])
		
			continue
		}
		if c=='\'' && b%2==0{
			
			quotes = append(quotes, str[cou:i] )
			cou=i+1
			b++
		}else if c=='\'' && b%2!=0 {
			
			quotes = append(quotes, "'"+strings.TrimSpace(str[cou:i])+"'" )
			cou=i+1
			b++
		}
	}
	str=SliceToString(quotes)
	quotes=Clean(str)
	return quotes
}