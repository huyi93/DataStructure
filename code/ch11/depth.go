//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnd69e/
//
//二叉树的最大深度
//给定一个二叉树，找出其最大深度。
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明:叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度3 。
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//执行用时： 4 ms , 在所有 Go 提交中击败了 89.77% 的用户
//内存消耗： 4.2 MB , 在所有 Go 提交中击败了 95.34% 的用户
func myMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	left, right := 0, 0
	if root.Left != nil {
		left += myMaxDepth(root.Left)
	}
	if root.Right != nil {
		right += myMaxDepth(root.Right)
	}
	return depth + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//执行用时： 4 ms , 在所有 Go 提交中击败了 89.77% 的用户
//内存消耗： 4.2 MB , 在所有 Go 提交中击败了 97.52% 的用户
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

func main() {
	t32 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	t31 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	t22 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	t21 := TreeNode{
		Val:   2,
		Left:  &t31,
		Right: &t32,
	}
	t := TreeNode{
		Val:   1,
		Left:  &t21,
		Right: &t22,
	}
	fmt.Println(maxDepth(&t))
	fmt.Println(maxDepth(&t22))
	fmt.Println(maxDepth(nil))
}
