package main

import "fmt"

//执行用时： 36 ms , 在所有 Go 提交中击败了 16.00% 的用户
//内存消耗： 8.2 MB , 在所有 Go 提交中击败了 5.02% 的用户
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, i int, j int) {
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

//执行用时： 32 ms , 在所有 Go 提交中击败了 17.04% 的用户
//内存消耗： 8.1 MB , 在所有 Go 提交中击败了 5.02% 的用户
func rotate2(nums []int, k int) {
	length := len(nums)
	k %= length
	tempSlice := make([]int, length)
	for i, v := range nums {
		tempSlice[(i+k)%length] = v
	}
	copy(nums, tempSlice)
}

func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 3)
	fmt.Println(nums)
}
