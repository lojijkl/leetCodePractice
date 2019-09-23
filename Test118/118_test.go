package Test118

import "testing"

func Test(t *testing.T) {
	t.Log(generate(5))
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	} else {
		res := [][]int{}
		temp := []int{}
		for i := 0; i < numRows; i++ {
			temp = printRows(temp)
			res = append(res, temp)
		}
		return res
	}
}

func printRows(data []int) []int {
	if len(data) == 0 {
		return []int{1}
	} else if len(data) == 1 {
		return []int{1, 1}
	}
	res := []int{1}
	for i := 1; i < len(data); i++ {
		res = append(res, data[i]+data[i-1])
	}
	res = append(res, 1)
	return res
}
