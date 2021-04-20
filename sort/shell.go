package sort

//插入排序
func ShellSort(s []int) []int {
	n := len(s)
	j := 0
	for d := n / 2; d >= 1; d = d / 2 {
		for i := d; i < n; i++ {
			if s[i] < s[i-d] {
				temp := s[i]
				for j = i - d; j >= 0 && temp < s[j]; j -= d {
					s[j+d] = s[j]
				}
				s[j+d] = temp
			}
		}
	}
	return s
}
