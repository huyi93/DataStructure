package sort

import (
	"fmt"
)

// 插入排序
func InsertSort(s []int) []int {
	j := 0
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			temp := s[i]
			for j = i - 1; j >= 0 && s[j] > temp; j-- {
				s[j+1] = s[j]
			}
			s[j+1] = temp
		}
	}
	return s
}

// 插入排序, 倒叙
func InsertSortDesc(s []int) []int {
	j := 0
	for i := 1; i < len(s); i++ {
		if s[i] > s[i-1] {
			temp := s[i]
			for j = i - 1; j >= 0 && s[j] < temp; j-- {
				s[j+1] = s[j]
			}
			s[j+1] = temp
		}
	}
	fmt.Println(s)
	return s
}

func InsertSortBySentry(s []int) []int {
	j := 0
	for i := 2; i < len(s); i++ {
		if s[i] < s[i-1] {
			s[0] = s[i]
			for j = i - 1; s[0] < s[j]; j-- {
				s[j+1] = s[j]
			}
			s[j+1] = s[0]
		}
	}
	fmt.Println(s)
	return s
}
