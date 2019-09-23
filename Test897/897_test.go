package Test897

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test(t *testing.T) {
	a := []int{5, 3, 6, 2, 4, 0, 8, 1, 0, 0, 0, 7}

	tree := &TreeNode{}
	tree.Val = a[0]
	tree.Left = &TreeNode{Val: 3}
	tree.Right = &TreeNode{Val: 6}
	tree.Left.Left = &TreeNode{Val: 2}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Left.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right.Right = &TreeNode{Val: 8}
	tree.Left.Right.Right.Left = &TreeNode{Val: 1}
	tree.Left.Right.Right.Right = &TreeNode{Val: 7}

	increasingBST(tree)
	t.Log()
}

func increasingBST(root *TreeNode) *TreeNode {
	tree := &TreeNode{}
	for root != nil {
		tree = root.Left
	}

}

func (tree *TreeNode) addToTree(data int) {
	if data < tree.Val && data != 0 {
		tree.Left.addToTree(data)
	} else if data > tree.Val {
		tree.Right.addToTree(data)
	}
}
