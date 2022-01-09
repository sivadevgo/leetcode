package main

//001 - WORKING
func sortedSquares(nums []int) []int {
	output := make([]int, len(nums))
	var j, v int
	for j, v = range nums {
		if v >= 0 {
			break
		}
	}

	var i int = j - 1
	var idx int

	for i >= 0 || j < len(nums) {
		if i < 0 {
			output[idx] = nums[j] * nums[j]
			idx++
			j++
		} else if j >= len(nums) {
			output[idx] = nums[i] * nums[i]
			idx++
			i--
		} else {
			if absolute(nums[i]) > nums[j] {
				output[idx] = nums[j] * nums[j]
				idx++
				j++
			} else {
				output[idx] = nums[i] * nums[i]
				idx++
				i--
			}
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

// 002 - WORKING

func sortedSquares(nums []int) []int {

	newNums := make([]int, len(nums))

	var rightIndex int
	var midValue int
	for rightIndex, midValue = range nums {
		if midValue >= 0 {
			break
		}
	}

	leftIndex := rightIndex - 1

	for i := 0; i < len(newNums); i++ {
		//Only rightIndex Valid
		if leftIndex < 0 {
			newNums[i] = nums[rightIndex] * nums[rightIndex]
			rightIndex++
		} else if rightIndex >= len(nums) {
			//Only leftIndex valid
			newNums[i] = nums[leftIndex] * nums[leftIndex]
			leftIndex--
		} else {
			if nums[rightIndex] < (-nums[leftIndex]) {
				newNums[i] = nums[rightIndex] * nums[rightIndex]
				rightIndex++
			} else {
				newNums[i] = nums[leftIndex] * nums[leftIndex]
				leftIndex--
			}
		}
	}
	return newNums

}
