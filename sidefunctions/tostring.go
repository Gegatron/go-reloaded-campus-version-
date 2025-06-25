package sidefunctions

func SliceToString(sli []string)string{
    str:=""
    for i,v:=range sli{
     
        if v!="" && i!=len(sli)-1 &&v[len(v)-1]!='\n' {
            str=str+v+" "
            continue
        }
        str=str+v
    }
    return str
}