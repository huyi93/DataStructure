//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
//
//两个数组的交集 II
//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1：
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//
//示例 2:
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]
//
//说明：
//输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
//我们可以不考虑输出结果的顺序。

//进阶：
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果nums1的大小比nums2小很多，哪种方法更优？
//如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		intersect(nums2, nums1)
	}

	dict := make(map[int]int)
	var resultInts []int
	for _, v := range nums1 {
		dict[v]++
	}
	for _, v := range nums2 {
		if dict[v] > 0 {
			resultInts = append(resultInts, v)
			dict[v]--
		}
	}

	return resultInts
}

//执行用时： 4 ms , 在所有 Go 提交中击败了 87.23% 的用户
//内存消耗： 3.1 MB , 在所有 Go 提交中击败了 26.32% 的用户
func myIntersect(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	var resultInts []int
	for _, v := range nums1 {
		if _, ok := dict[v]; ok {
			dict[v]++
		} else {
			dict[v] = 1
		}
	}
	for _, v := range nums2 {
		if value, ok := dict[v]; ok {
			if value != 0 {
				resultInts = append(resultInts, v)
				dict[v]--
			}
		}
	}

	return resultInts
}

//执行用时： 4 ms , 在所有 Go 提交中击败了 87.23% 的用户
//内存消耗： 2.8 MB , 在所有 Go 提交中击败了 83.20% 的用户
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	a := 0
	b := 0
	var resultInts []int
	length1, length2 := len(nums1), len(nums2)

	for length1 > a && length2 > b {
		if nums1[a] < nums2[b] {
			a++
		} else if nums1[a] == nums2[b] {
			resultInts = append(resultInts, nums1[a])
			a++
			b++
		} else {
			b++
		}
	}

	return resultInts
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
