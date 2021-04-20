package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {

	if len(S) == 0 {
		return ""
	}

	str := ""
	cnt := 1
	i := 0
	j := 1

	for ; j < len(S); j++ {
		if S[i:i+1] == S[j:j+1] {
			cnt++
		} else {
			str = str + S[i:i+1] + strconv.Itoa(cnt)
			i = j
			cnt = 1
		}
	}
	str = str + S[i:i+1] + strconv.Itoa(cnt)

	if len(str) < len(S) {
		return str
	}

	return S
}

//执行用时：88 ms , 在所有 Go 提交中击败了 38.77%
//内存消耗：8.3 MB , 在所有 Go 提交中击败了 32.97%
func myCompressString(S string) string {

	if len(S) <= 2 {
		return S
	}

	str := int32(S[0])
	cnt := 0
	ans := ""

	for _, ch := range S {
		if ch == str {
			cnt += 1
		} else {
			ans += string(str) + strconv.Itoa(cnt)
			cnt = 1
			str = ch
		}
	}
	ans += string(str) + strconv.Itoa(cnt)
	if len(ans) < len(S) {
		return ans
	}
	return S
}

func main() {
	s := ""
	fmt.Println(myCompressString(s))
}
