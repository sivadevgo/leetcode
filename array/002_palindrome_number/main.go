func isPalindrome(x int) bool {
    if x<0{
        return false
    }else if x<10{
        return true
    }
    
    input := fmt.Sprintf("%d",x)
    n := len(input)
    for i:=0;i<n/2;i++{
        if input[i]!=input[n-1-i]{
            return false
        }
    }
    
    return true
}


