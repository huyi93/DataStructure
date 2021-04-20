//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnarn7/
//
//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。
//
//现有一个链表 --head =[4,5,1,9]
//
//示例 1：
//输入：head = [4,5,1,9], node = 5
//输出：[4,1,9]
//解释：给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//
//示例 2：
//
//输入：head = [4,5,1,9], node = 1
//输出：[4,5,9]
//解释：给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	d := ListNode{3, nil}
	c := ListNode{1, &d}
	b := ListNode{5, &c}
	a := ListNode{4, &b}
	fmt.Println(a.Next.Next.Val)
	deleteNode(&c)
	fmt.Println(a.Next.Next.Val)
}
