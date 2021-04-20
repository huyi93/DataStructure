package tree

import "fmt"

//二叉树
type BinaryTree struct {
	Value       interface{}
	Left, Right *BinaryTree
}

//线索二叉树结点
type ThreadNode struct {
	Value       interface{}
	Left, Right *BinaryTree
	LTag, RTag  int // 0 指向 孩子 1 指向线索
}

func Init() *BinaryTree {
	return new(BinaryTree)
}

func (t *BinaryTree) Clear() {
	if t.Value != nil {
		t.Value = nil
		t.Left = nil
		t.Right = nil
	}
}

func (t *BinaryTree) IsEmpty() bool {
	if t.Value == nil {
		return true
	}
	return false
}

func PreOrder(t *BinaryTree) {
	if t.Value != nil {
		visit(t)
		PreOrder(t.Left)
		PreOrder(t.Right)
	}
}

func InOrder(t *BinaryTree) {
	if t.Value != nil {
		InOrder(t.Left)
		visit(t)
		InOrder(t.Right)
	}
}

func PostOrder(t *BinaryTree) {
	if t.Value != nil {
		PostOrder(t.Left)
		PostOrder(t.Right)
		visit(t)
	}
}

func visit(t *BinaryTree) {
	fmt.Println(t.Value, t.Left, t.Right)
}
