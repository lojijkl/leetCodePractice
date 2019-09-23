package Test965

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test(t *testing.T) {
	data := []int{1, 1, 1, 1, 1, nil, 1}
	t.Log(isUnivalTree(data))
}
func isUnivalTree(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Val != root.Left.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
