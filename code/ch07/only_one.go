//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
//
//只出现一次的数字
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//输入: [2,2,1]
//输出: 1
//
//示例2:
//输入: [4,1,2,1,2]
//输出: 4

//交换律：a ^ b ^ c <=> a ^ c ^ b
//
//任何数于0异或为任何数 0 ^ n => n
//
//相同的数异或为0: n ^ n => 0
package main

import (
	"fmt"
	"sort"
)

//执行用时： 32 ms , 在所有 Go 提交中击败了 9.08% 的用户
//内存消耗： 6.1 MB , 在所有 Go 提交中击败了 74.90% 的用户
func mySingleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			return nums[i]
		} else {
			for nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return nums[len(nums)-1]
}

//异或
//执行用时： 20 ms , 在所有 Go 提交中击败了 77.98% 的用户
//内存消耗： 6.1 MB , 在所有 Go 提交中击败了 88.92% 的用户
func singleNumber(nums []int) int {
	a := 0
	for _, num := range nums {
		a = a ^ num
	}
	return a
}

func main() {
	list := []int{4, 1, 2, 1, 2, 5, 4, 1}
	fmt.Println(singleNumber(list))
}
