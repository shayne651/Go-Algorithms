package binary

func BinarySearch(arr []int8, find int8) bool {
	low := 0
	high := len(arr)
	for  {
		midIndex := low + ((high-low)/2)
		midVal := arr[midIndex]
		if midVal == find{
			return true
		} else if midVal > find {
			high = midIndex - 1
		} else {
			low = midIndex
		}
	}
	return false
}
