func findMaxConsecutiveOnes(nums []int) int {
    var maxOne int
    var currentOne int
    for _,element := range nums{
        if element!=1{
            if maxOne<currentOne{
                maxOne = currentOne
            }
            currentOne = 0            
        }else{
            currentOne++
        }        
    }
            if maxOne<currentOne{
                maxOne = currentOne
            }
    return maxOne
    
}
