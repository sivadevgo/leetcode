package main

import "fmt"

//1. WORKING
// func removeElement(nums []int, val int) int {
//     l := len(nums)
//     for i:=0;i<l;i++{
//         if nums[i]==val{
//             moveLeft(nums,i)
//             l--
//             i--
//         }
//     }
//     return l
// }

// func moveLeft(nums []int,index int){
//     for i:=index;i<len(nums)-1;i++{
//         nums[i] = nums[i+1]
//     }
// }

//2. WORKING
// func removeElement(nums []int, val int) int {
//     for i:=0;i<len(nums);i++{
//         if nums[i]==val{
//             nums = append(nums[:i],nums[i+1:]...)
//             i--
//         }

//     }
//     return len(nums)
// }

//3. NOT WORKING
// func removeElement(nums []int, val int) int {

//     var rIndex = []int{}

//     for i,v := range nums{
//         if val==v{
//             rIndex = append(rIndex,i)
//         }
//     }

//     l := len(nums)
//     var r int
//     var j int
//     for i:=0;i<l;i++{
//         if i==rIndex[j]-r{
//             j++
//             if i+1+r<len(nums){
//                 nums[i]=nums[i+1]
//             }
//             r++
//             i--
//             l--
//         }
//     }
//     return l
// }

//4. WORKING
// func removeElement(nums []int, val int) int {
// 	l := 0
// 	r := len(nums) - 1

// 	for l <= r {
// 		if nums[l] == val {
// 			nums[l] = nums[r]
// 			r--
// 		} else {
// 			l++
// 		}
// 	}

// 	return l
// }

//5. WORKING
func removeElement(nums []int, val int) int {
	var i int

	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

func main() {
	// s := []int{1, 2, 3, 2}
	v := 2
	s := []int{2, 3, 2}
	o := removeElement(s, v)
	fmt.Println(o)
	fmt.Println(s)
}
