package binary

//func BinarySearch(arr []int8, find int8) bool {
//
//	half := len(arr)/2
//	if half == 0 {
//		return false
//	} else if half == 1 && arr[0] == find{
//		return true
//	} else if half == 1 && arr[0] != find{
//		return false
//	}
//
//	mid := arr[half]	
//
//	if mid == find {
//		return true
//	} else if mid > find {
//		return BinarySearch(arr[0:half + 1], find)
//	} else {
//		return BinarySearch(arr[half - 1:len(arr)], find)
//	}
//	return false
//}

func BinarySearch(arr []int8, find int8) bool {
	index := len(arr)/2
	if len(arr) > 0 {
		val := arr[index]
		if val == find {
			return true
		} else if len(arr) == 1 {
			return false
		} else if val > find {
			return BinarySearch(arr[0:index], find) 
		} else {
			return BinarySearch(arr[index:len(arr)], find)
		}
	}
	return false
}
