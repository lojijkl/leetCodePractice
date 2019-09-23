package Test349

import "testing"

func Test(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	t.Log(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	temp := map[int]bool{}
	for _, v := range nums1 {
		temp[v] = true
	}
	res := []int{}
	for _, v := range nums2 {
		if temp[v] == true {
			res = append(res, v)
			temp[v] = false
		}
	}
	return res
}
