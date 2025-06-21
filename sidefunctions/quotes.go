package sidefunctions

//import "strings"

func Quotes(reloaded []string)[]string{
	var quotes []string

	//cou:=0
	b:=false
	for i := 0; i <len(reloaded); i++ {
		if reloaded[i]!="" && !b && reloaded[i][len(reloaded[i])-1]=='\'' && i+1<len(reloaded) {
			if reloaded[i+1]!="'" {
				reloaded[i+1]=" "+"'"+reloaded[i+1]
			reloaded[i]=reloaded[i][:len(reloaded[i])-1]
			b=true
			}
			
		}else if reloaded[i]!="" && b && reloaded[i][0]=='\'' && i>0{
			if reloaded[i-1]!="'" {
				reloaded[i-1]=reloaded[i-1]+"'"
			reloaded[i]=reloaded[i][1:]
			b=false
			}
			
	}
}
	str:=SliceToString(reloaded)
	quotes=Clean(str)
	
	return quotes
}