package search

//二分法查找, 升序排列
func BinarySearch(s []int, key int) int {
	low, high, mid := 0, len(s), 0
	if high == 0 || s[low] > key || s[high-1] < key {
		return -1
	}
	for low <= high {
		mid = (low + high) / 2
		if s[mid] == key {
			return mid
		}
		if s[low] == key {
			return low
		}
		if s[high] == key {
			return low
		}

		if s[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}

func BinarySearchDesc(s []int, key int) int {
	low, high, mid := 0, len(s), 0
	if high == 0 {
		return -1
	}
	for low <= high {
		mid = (low + high) / 2
		if s[mid] == key {
			return mid
		} else if s[mid] < key {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}
