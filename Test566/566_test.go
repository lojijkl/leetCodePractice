package Test566

import (
	"testing"
)

//重塑矩阵

func Test(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}
	r, c := 1, 4
	res := matrixReshape(nums, r, c)
	t.Log(res)
}

func Benchmark_test(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) < r*c {
		return nums
	}
	data := make([][]int, r)
	var cols int
	d := make([]int, c)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			d[cols] = nums[i][j]
			cols++
			if cols == c { //要返回的数据进行下一行遍历
				data = append(data, d)
				d = make([]int, c)
				cols = 0
			}
		}
	}
	return data
}
