package Test867

import "testing"

func Test(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	b := transpose(a)
	t.Log(b)
}

func transpose(A [][]int) [][]int {
	res := make([][]int, len(A[0]))

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			res[j] = append(res[j], A[i][j])
		}
	}
	return res
}
