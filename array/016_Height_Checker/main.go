package main

import "sort"

func heightChecker(heights []int) int {
	var output int
	expected := make([]int, len(heights))
	for i, val := range heights {
		expected[i] = val
	}

	sort.Ints(expected)

	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			output++
		}
	}

	return output
}
