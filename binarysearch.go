package binarysearch

func BinarySearch(s []int, key int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		middleIndex := (low + high) >> 1
		middleValue := s[middleIndex]

		if middleValue < key {
			low = middleIndex + 1
		} else if middleIndex > key {
			high = middleIndex - 1
		} else {
			return middleIndex
		}
	}
	return -(low + 1)
}

func BinarySearchDiv(s []int, key int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		middleIndex := (low + high) / 2
		middleValue := s[middleIndex]

		if middleValue < key {
			low = middleIndex + 1
		} else if middleIndex > key {
			high = middleIndex - 1
		} else {
			return middleIndex
		}
	}
	return -(low + 1)
}
