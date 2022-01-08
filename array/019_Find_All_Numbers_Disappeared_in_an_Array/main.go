package main

func findDisappearedNumbers(nums []int) []int {
	output := []int{}
	for _, v := range nums {
		var idx int
		if v < 0 {
			idx = absolute(v) - 1
		} else {
			idx = v - 1
		}

		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i, v := range nums {
		if v > 0 {
			output = append(output, i+1)
		}
	}

	return output
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
