package Test1122

import (
	"sort"
	"testing"
)

func Test(t *testing.T) {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	data := relativeSortArray(arr1, arr2)
	t.Log(data)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	maps := map[int]int{}
	for _, v := range arr1 {
		maps[v]++
	}
	res := []int{}
	for _, v := range arr2 {
		for i := 0; i < maps[v]; i++ {
			res = append(res, v)

		}
		delete(maps, v)
	}

	other := []int{}
	for k, v := range maps {
		for i := 0; i < v; i++ {
			other = append(other, k)
		}
	}
	sort.Ints(other)
	for _, v := range other {
		res = append(res, v)
	}
	return res
}
