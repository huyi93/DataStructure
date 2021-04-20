//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn2925/
//
//删除链表的倒数第N个节点
//给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
//
//进阶：你能尝试使用一趟扫描实现吗？
//
//示例 1：
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//示例 2：
//输入：head = [1], n = 1
//输出：[]
//
//示例 3：
//输入：head = [1,2], n = 1
//输出：[1]
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗： 2.2 MB , 在所有 Go 提交中击败了 100.00% 的用户
func myRemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var num int
	list := make([]*ListNode, 30)

	for node := head.Next; node != nil; node = node.Next {
		num++
		list[num] = node
	}

	if n != num+1 {
		list[num-n].Next = list[num-n].Next.Next
		return list[0]

	} else {
		return head.Next
	}
}

func myRemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	var list []*ListNode
	dummy := &ListNode{0, head}

	for node := dummy; node != nil; node = node.Next {
		list = append(list, node)
	}

	list[len(list)-n-1].Next = list[len(list)-n-1].Next.Next
	return dummy.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

//方法一：计算链表长度
//执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗： 2.2 MB , 在所有 Go 提交中击败了 86.97% 的用户
//复杂度分析
//时间复杂度：O(L)，其中 L 是链表的长度。
//空间复杂度：O(1)。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

func myRemoveNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for node := head; node != nil; node = node.Next {
		first = first.Next
		if n > 0 {
			n--
		} else {
			second = second.Next
		}
	}
	second.Next = second.Next.Next

	return dummy.Next
}

//双指针
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next

}

func main() {
	g := ListNode{Val: 5}
	f := ListNode{Val: 4, Next: &g}
	e := ListNode{Val: 5}
	d := ListNode{Val: 4, Next: &e}
	c := ListNode{Val: 3, Next: &d}
	b := ListNode{Val: 2, Next: &c}
	a := ListNode{Val: 1, Next: &b}
	fmt.Println(a.Next.Next.Val)
	removeNthFromEnd3(&a, 3)
	fmt.Println(a.Next.Next.Val)
	removeNthFromEnd3(&f, 2)
	fmt.Println(f.Next)
	fmt.Println(f)
}
