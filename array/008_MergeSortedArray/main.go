func merge(nums1 []int, m int, nums2 []int, n int)  {

    var i,j int
    for i,j = 0,0;i<n&&j<m+n;{
        if nums1[j]<nums2[i]{
            j++
        }else{
            moveRight(nums1,j)
            nums1[j]=nums2[i]
            i++
            j++
        }
    }
    
    r := n-i
    for ;i<n;i++{
        nums1[len(nums1)-r] = nums2[i]
        r--
    }
}

func moveRight(arr []int,index int){
    for i:= len(arr)-2;i>=index;i--{
        arr[i+1]=arr[i]
    }
}
