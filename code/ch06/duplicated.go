//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
//
//存在重复元素
//给定一个整数数组，判断是否存在重复元素。
//如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
//
//示例 1:
//输入: [1,2,3,1]
//输出: true
//
//示例 2:
//输入: [1,2,3,4]
//输出: false
//
//示例3:
//输入: [1,1,1,3,3,4,3,2,4,2]
//输出: true
package main

import (
	"fmt"
	"sort"
)

func myContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	list := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(myContainsDuplicate(list))
}
