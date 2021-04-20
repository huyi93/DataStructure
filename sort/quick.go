package sort

//快排
func QuickSort(s []int, low int, high int) {
	//low, high := 0, len(s)
	//pivot := rand.Intn(high)
	if low < high {
		pivot := partition(s, low, high)
		QuickSort(s, low, pivot-1)
		QuickSort(s, pivot+1, high)
	}

}

func partition(s []int, low int, high int) int {
	pivot := s[low]
	for low < high {
		for ; low < high && s[high] >= pivot; high-- {
		}
		s[low] = s[high]
		for ; low < high && s[low] <= pivot; low++ {
		}
		s[high] = s[low]

	}
	s[low] = pivot
	return low
}
