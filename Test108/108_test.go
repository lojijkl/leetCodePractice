package Test108

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	data := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(data))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums) - 1
	if l < 0 {
		return nil
	}
	mid := (l) >> 1
	tree := &TreeNode{
		Val:   nums[mid],
		Right: sortedArrayToBST(nums[mid+1:]),
		Left:  sortedArrayToBST(nums[:mid]),
	}
	return tree
}
