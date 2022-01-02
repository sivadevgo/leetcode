// func findNumbers(nums []int) int {
//     var output int
//     for _,element := range nums{
//         if isEven(element){
//             output++
//         }
//     }
//     return output
    
// }


// func isEven(input int) bool{
//     var output int
//     for input!=0{
//         input = input/10
//         output++
//     }
//     return output%2==0
// }


func findNumbers(nums []int) int {
    
    var output int
    for _,element := range nums{
        if len(fmt.Sprintf("%d",element))%2==0{
            output++
        }
    }
    return output
    
}
