package sort

//冒泡排序
func BubbleSort(s []int) []int {
	length := len(s)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
