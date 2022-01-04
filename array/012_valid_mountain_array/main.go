//001 WORKING
func validMountainArray(arr []int) bool {
	var incC int
	var decC int
	var raise bool
	for i := 0; i < len(arr)-1; i++ {
		if incC < decC {
			return false
		}
		if arr[i] == arr[i+1] {
			return false
		}

		if incC == 0 && !raise && arr[i] > arr[i+1] {
			return false
		}

		if !raise && arr[i] < arr[i+1] {
			incC++
			raise = true
		} else if raise && arr[i] > arr[i+1] {
			decC++
			raise = false
		}
	}
	return incC == 1 && decC == 1
}