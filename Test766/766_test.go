package Test766

import (
	"testing"
)

func Test(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	t.Log(isToeplitzMatrix(matrix))
}

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) < 1 {
		return false
	}
	temp1 := []int{}
	temp2 := []int{}
	for i := 0; i < len(matrix)-1; i++ {
		temp1 = matrix[i]
		temp2 = matrix[i+1]
		for j := 0; j < len(matrix[0])-1; j++ {
			if temp1[j] != temp2[j+1] {
				return false
			}
		}
	}

	return true
}
