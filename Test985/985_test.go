package Test985

import (
	"testing"
)

func Test(t *testing.T) {
	A := []int{1, 2, 3, 4}
	queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	res := sumEvenAfterQueries(A, queries)
	t.Log(res)
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := []int{}
	sum := 0
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}
	for _, query := range queries {
		old := A[query[1]]
		tempRes := old + query[0]
		if old%2 == 0 && tempRes%2 == 0 {
			sum = sum - old + tempRes
		} else if old%2 == 0 && tempRes%2 != 0 {
			sum = sum - old
		} else if old%2 != 0 && tempRes%2 == 0 {
			sum = sum + tempRes
		}
		res = append(res, sum)
		A[query[1]] = tempRes
	}
	return res
}
