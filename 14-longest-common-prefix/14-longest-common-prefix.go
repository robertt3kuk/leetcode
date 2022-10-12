func longestCommonPrefix(strs []string) string {
    str:=""
    minlen:=200
    for _,c := range strs{
        if len(c)<minlen{
            minlen=len(c)
        }    
    }
    for i:=0;i<minlen;i++{
        if got(strs,i){
            str+=string(strs[0][i])
        }else{
            break
        }
    }
    return str
    
    

}
func got(strs []string,i int)bool{
    s:=strs[0][i]
    for _,c:= range strs{
        if s!=c[i]{
            return false
        }
    }
    return true
    
}