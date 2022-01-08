package main

import "sort"

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	sort.Ints(nums)

	output := make([]int, 3)
	var i int = 2
	output[i] = nums[len(nums)-1]
	i--
	for j := len(nums) - 2; j >= 0 && i >= 0; j-- {
		if nums[j] != nums[j+1] {
			output[i] = nums[j]
			i--
		}
	}

	if i == -1 {
		return output[0]
	} else {
		return output[2]
	}
}
