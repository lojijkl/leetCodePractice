package Test654

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test(t *testing.T) {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := maxIndex(nums)
	tree := &TreeNode{
		Val:   nums[index],
		Right: constructMaximumBinaryTree(nums[index+1:]),
		Left:  constructMaximumBinaryTree(nums[:index]),
	}
	return tree
}

//找出数组中最大的下标
func maxIndex(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var max = 0
	for i := 1; i < len(nums); i++ {
		if nums[max] < nums[i] {
			max = i
		}
	}
	return max
}
