//链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn08xg/
//
//验证二叉搜索树
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例1:
//
//输入:
//    2
//   / \
//  1   3
//输出: true
//示例2:
//
//输入:
//    5
//   / \
//  1   4
//    / \
//   3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//    根节点的值为 5 ，但是其右子节点值为 4 。
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root.Right == nil && root.Left == nil {
		return true
	}

	if root.Left != nil {
		isValidBST(root.Left)
	}
	if root.Right != nil {
		isValidBST(root.Right)
	}

	if root.Left != nil && root.Val < root.Right.Val {
		return true
	}
	if root.Right != nil && root.Val < root.Right.Val {
		return true
	}

	return false
}

func isTure(a, b bool) bool {
	if a != true || b != true {
		return false
	}
	return true
}

func main() {
	t32 := TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	t31 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	t22 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	t21 := TreeNode{
		Val:   4,
		Left:  &t31,
		Right: &t32,
	}
	t := TreeNode{
		Val:   5,
		Left:  &t22,
		Right: &t21,
	}
	fmt.Println(isValidBST(&t))
}
