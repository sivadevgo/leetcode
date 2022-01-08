package main

import "fmt"

//001 WORKING
// func validMountainArray(arr []int) bool {
// 	var incC int
// 	var decC int
// 	var raise bool
// 	for i := 0; i < len(arr)-1; i++ {
// 		if incC < decC {
// 			return false
// 		}
// 		if arr[i] == arr[i+1] {
// 			return false
// 		}

// 		if incC == 0 && !raise && arr[i] > arr[i+1] {
// 			return false
// 		}

// 		if !raise && arr[i] < arr[i+1] {
// 			incC++
// 			raise = true
// 		} else if raise && arr[i] > arr[i+1] {
// 			decC++
// 			raise = false
// 		}
// 	}
// 	return incC == 1 && decC == 1
// }

//002 - Working

// func validMountainArray(arr []int) bool {
// 	//lenth should be >=3
// 	if len(arr) < 3 {
// 		return false
// 	}

// 	//find peak
// 	var peak int
// 	for peak = 0; peak < len(arr)-1; peak++ {
// 		if arr[peak] >= arr[peak+1] {
// 			break
// 		}
// 	}

// 	if peak+1 == len(arr) || peak == 0 {
// 		return false
// 	}

// 	//check from both ends till peak
// 	var i int
// 	var j int = len(arr) - 1
// 	for i < peak || j > peak {
// 		if i != peak && arr[i] >= arr[i+1] {
// 			return false
// 		} else if i != peak {
// 			i++
// 		}

// 		if j != peak && arr[j] >= arr[j-1] {
// 			return false
// 		} else if j != peak {
// 			j--
// 		}
// 	}

// 	if i == peak && j == peak {
// 		return true
// 	} else {
// 		return false
// 	}
// }

//003 - WORKING
// func validMountainArray(arr []int) bool {
// 	//lenth should be >=3
// 	if len(arr) < 3 {
// 		return false
// 	}

// 	//find peak
// 	var peak int
// 	for peak = 0; peak < len(arr)-1; peak++ {
// 		if arr[peak] >= arr[peak+1] {
// 			break
// 		}
// 	}

//     var peakR int = len(arr)-1

//     for peakR = len(arr)-1;peakR>=1;peakR--{
//         if arr[peakR]>=arr[peakR-1]{
//             break
//         }
//     }

//     return peak==peakR && peak!=0 && peakR!=len(arr)-1

// }

//004
// func validMountainArray(arr []int) bool {
// 	//lenth should be >=3
// 	if len(arr) < 3 {
// 		return false
// 	}

// 	//find peak
// 	var peak int
// 	for peak = 0; peak < len(arr)-1; peak++ {
// 		if arr[peak] >= arr[peak+1] {
// 			break
// 		}
// 	}

// 	var peakR int = len(arr) - 1

// 	for peakR = len(arr) - 1; peakR >= 1; peakR-- {
// 		if arr[peakR] >= arr[peakR-1] {
// 			break
// 		}
// 	}

// 	return peak == peakR && peak != 0 && peakR != len(arr)-1

// }

//005

func validMountainArray(arr []int) bool {
	//lenth should be >=3
	if len(arr) < 3 {
		return false
	}

	//find peak
	var peak int
	for peak = 0; peak < len(arr)-1; peak++ {
		if arr[peak] >= arr[peak+1] {
			break
		}
	}

	if peak == len(arr)-1 || peak == 0 {
		return false
	}

	for ; peak < len(arr)-1; peak++ {
		if arr[peak] <= arr[peak+1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr))
}
