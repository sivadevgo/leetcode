//001 WORKING

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == 2*arr[j] || 2*arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

//002 WORKING
func checkIfExist(arr []int) bool {
	m := map[int]bool{}

	for _, v := range arr {
		if v%2 == 0 && m[v/2] || m[v*2] {
			return true
		}
		m[v] = true
	}
	return false
}

//003_WORKING
func checkIfExist(arr []int) bool {
	m := map[float64]bool{}

	for _, v := range arr {
		f := float64(v)
		if m[f/2.0] || m[f*2.0] {
			return true
		}
		m[f] = true
	}
	return false
}

