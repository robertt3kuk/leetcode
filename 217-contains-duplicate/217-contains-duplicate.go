func containsDuplicate(nums []int) bool {
    n:=make(map[int]int)
    for _,c := range nums{
        _,ok:=n[c]
        if !ok{
            n[int(c)]=int(c)
        }else{
            return true
        }
        
    }
    
   return false
}