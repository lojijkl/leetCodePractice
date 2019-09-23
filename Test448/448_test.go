package Test448

import (
	"sort"
	"testing"
)

func Test(t *testing.T) {
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	b := findDisappearedNumbers(a)
	t.Log(b)
}

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		for nums[i] > nums[i-1]+1 {
			nums[i-1] = nums[i-1] + 1
			res = append(res, nums[i-1])
		}
	}
	return res
}
