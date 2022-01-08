package main

// func moveZeroes(nums []int)  {
//     var i int

//     for j:=0;j<len(nums);j++{
//         if nums[j]!=0{
//             nums[i],nums[j]=nums[j],nums[i]
//             i++
//         }
//     }
// }

func moveZeroes(nums []int) {
	var i int

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
