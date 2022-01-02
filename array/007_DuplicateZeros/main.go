func duplicateZeros(arr []int)  {
    for i:=0;i<len(arr)-1;i++{
        if arr[i]==0{
            moveRight(arr,i)  
            i++
        }
    }
}


func moveRight(arr []int,index int){
    for i:=len(arr)-2;i>=index;i--{
        arr[i+1] = arr[i]
    }    
}
