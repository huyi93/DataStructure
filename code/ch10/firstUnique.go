//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r
//字符串中的第一个唯一字符
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//示例：
//s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
//
package main

import (
	"fmt"
	"strings"
)

//执行用时： 36 ms , 在所有 Go 提交中击败了 45.20% 的用户
//内存消耗： 5.4 MB , 在所有 Go 提交中击败了 26.41% 的用户
func firstUniqChar(s string) int {
	dict := map[int32]int{}
	for _, i := range s {
		dict[i]++
	}
	for i, v := range s {
		if dict[v] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {

	for i, v := range s {
		if strings.IndexAny(s, string(v)) == strings.LastIndexAny(s, string(v)) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar2("loveleetcode"))
}
