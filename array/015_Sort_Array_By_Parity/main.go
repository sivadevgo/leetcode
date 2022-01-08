package main

func sortArrayByParity(nums []int) []int {
	var i int
	for j := 0; j < len(nums); j++ {
		for i < len(nums) && nums[i]%2 == 0 {
			i++
		}
		if j < i {
			j = i
			continue
		}
		if nums[j]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
