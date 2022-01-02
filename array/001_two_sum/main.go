func twoSum(nums []int, target int) []int {
    
    m := make(map[int]int)
       
    for i := 0;i<len(nums);i++{
        if idx,ok := m[target-nums[i]];ok{
            return []int{idx,i}
        }
        m[nums[i]]=i        
    }    
    return nil
}


