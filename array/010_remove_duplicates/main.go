func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return len(nums)
	}
	var j int

	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]

		}
	}
	return j + 1
}