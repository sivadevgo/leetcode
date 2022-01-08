package main

import (
	"fmt"
)

func replaceElements(arr []int) []int {
	t := arr[len(arr)-1]
	max := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		if t > max {
			max = t
		}
		t = arr[i]
		arr[i] = max
	}
	return arr
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(arr)
	fmt.Println(replaceElements(arr))
}
