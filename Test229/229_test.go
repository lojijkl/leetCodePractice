package Test229

import "testing"

func Test(t *testing.T) {
	a := []int{3, 2, 3}
	b := majorityElement(a)
	t.Log(b)
}

func majorityElement(nums []int) []int {
	maps := map[int]int{}
	l := len(nums) / 3
	for _, v := range nums {
		maps[v]++
	}
	res := []int{}
	for k, v := range maps {
		if v > l {
			res = append(res, k)
		}
	}
	return res
}
