func reverseString(s []byte)  {
    f:=0
    l:=len(s)-1
    for f<l{
        s[f],s[l]=s[l],s[f]
        f++
l--
    }
    
}