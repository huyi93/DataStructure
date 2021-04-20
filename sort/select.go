package sort

//选择排序
func SimpleSelectSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		j := selectMinKey(s[i+1:])
		//fmt.Println(j)
		j = j + i + 1
		if s[i] > s[j] {
			s[i], s[j] = s[j], s[i]
		}
	}
	return s
}

func selectMinKey(s []int) int {
	minKey, minValue := 0, s[0]
	for k, v := range s {
		if minValue > v {
			minKey = k
			minValue = v
		}
	}
	return minKey
}
