package Test4

import (
	"sort"
	"testing"
)

func Test(t *testing.T) {
	//t.Log(findMedianSortedArrays([]int{1,2},[]int{3,4}))
	//t.Log(findMedianSortedArrays([]int{0,0},[]int{0,0}))
	//t.Log(findMedianSortedArrays([]int{3},[]int{-2,-1}))
	//t.Log(findMedianSortedArrays([]int{2},[]int{}))
	t.Log(findMedianSortedArrays([]int{}, []int{2, 3}))
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l := len(nums1)
	if l < 1 {
		return 0
	} else if l%2 == 1 {
		return float64(nums1[l/2])
	} else {
		return float64((nums1[(l-1)/2] + nums1[(l)/2])) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	k := l >> 1
	i, j, mid, mid1 := 0, 0, 0, 0
	for i+j < k+1 {
		if i+j == k {
			mid1 = mid
		}
		if i > l1-1 {
			mid = nums2[j]
			j++
			continue
		}
		if j > l2-1 {
			mid = nums1[i]
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			mid = nums1[i]
			i++
		} else {
			mid = nums2[j]
			j++
		}

	}
	if l%2 == 1 {
		return float64(mid)
	} else {
		return float64(mid+mid1) / 2
	}
}

func getKkey(nums1, nums2 []int, l int) []int {
	nums := []int{}
	for len(nums) < l/2+1 && len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			nums = append(nums, nums1[0])
			nums1 = nums[1:]
		} else {
			nums = append(nums, nums2[0])
			nums2 = nums[1:]
		}
	}
	return nums
}

func abs(l1, l2 int) int {
	if l1 > l2 {
		return l1
	} else {
		return l2
	}
}
