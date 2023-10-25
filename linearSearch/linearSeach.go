package linear

func LinearSearch(arr []int8, find int8) bool {
	for n := 0; n < len(arr); n ++ {
		if arr[n] == find {
			return true
		}
	}
	return false
}
