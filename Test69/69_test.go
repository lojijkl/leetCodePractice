package Test69

import "testing"

func Test(t *testing.T) {
	t.Log(mySqrt(8))
}

func mySqrt(x int) int {
	max := x
	min := 0
	for min < max {
		data := (min + max + 1) / 2
		square := data * data
		if square > x {
			max = data - 1
		} else {
			min = data
		}
	}
	return min
}
