//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/
//
//反转字符串
//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
//
//示例 1：
//输入：['h','e','l','l','o']
//输出：['o','l','l','e','h']

//示例 2：
//输入：['H','a','n','n','a','h']
//输出：['h','a','n','n','a','H']
package main

import (
	"fmt"
)

//执行用时： 36  ms , 在所有 Go 提交中击败了 95.11% 的用户
//内存消耗： 6.5 MB , 在所有 Go 提交中击败了 97.02% 的用户
func myReverseString(s []byte) {
	length := len(s) - 1
	for i := 0; length > i; i++ {
		s[i], s[length] = s[length], s[i]
		length--
	}
}

func main() {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	fmt.Println(s)
	myReverseString(s)
	fmt.Println(s)
}
