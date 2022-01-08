package main

import "fmt"

func consecutive(nums []int) int {
	var count int
	var currentCount int

	for j := 0; j < len(nums); j++ {
		if nums[j] == 1 {
			currentCount++
		}
		if nums[j] != 1 || j == len(nums)-1 && nums[j] == 1 {
			if currentCount > count {
				count = currentCount
			}
			currentCount = 0
		}
	}
	return count
}

func main() {
	arr1 := []int{1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1}
	fmt.Println(consecutive(arr1))

	arr2 := []int{0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}
	fmt.Println(consecutive(arr2))

}
