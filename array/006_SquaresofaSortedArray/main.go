
func sortedSquares(nums []int) []int {
    
    newNums := make([]int,len(nums))
    
    var rightIndex int   
    var midValue int
    for rightIndex,midValue = range nums {
        if midValue >=0{
            break
        }
    }
    
    leftIndex := rightIndex - 1
    
    for i:=0;i<len(newNums);i++{
        //Only rightIndex Valid
        if leftIndex<0{
            newNums[i] = nums[rightIndex]*nums[rightIndex]
            rightIndex++
        } else if rightIndex>=len(nums){
            //Only leftIndex valid
            newNums[i] = nums[leftIndex]*nums[leftIndex]
            leftIndex--
        } else{
            if nums[rightIndex]<(-nums[leftIndex]){
                newNums[i] = nums[rightIndex]*nums[rightIndex]
                rightIndex++
            }else{
                newNums[i] = nums[leftIndex]*nums[leftIndex]
                leftIndex--
            }
        }
    }
    return newNums
    
}

