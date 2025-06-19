package sidefunctions

func Quotes(reloaded []string)[]string{
	b:=false
	for i,c:=range reloaded{
		 if c[0]=='\'' && b && i>0{
			reloaded[i-1]=reloaded[i-1]+"'"
			if len(c)==1 {
				reloaded[i]=""
				b=false
				continue
			}
			reloaded[i]=string(reloaded[i][0:])
			b=false
		}else if c[len(c)-1]=='\'' && !b  && i<len(reloaded)-1{
			reloaded[i+1]="'"+reloaded[i+1]
			if len(c)==1 {
				reloaded[i]=""
				b=true
				continue
			}
			b=true
			reloaded[i]=string(reloaded[i][:len(reloaded[i])-1])
		}
	}
	str:=SliceToString(reloaded)
	reloaded=Clean(str)
	return reloaded
}