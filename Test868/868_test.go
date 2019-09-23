package Test868

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(binaryGap(5))
	t.Log(binaryGap(6))
	t.Log(binaryGap(8))
}
func binaryGap(N int) int {
	var max, cnt int
	var start bool
	for N > 0 {
		if N^1 < N {
			start = true
			if cnt >= max {
				max = cnt
			}
			cnt = 1
		} else if start {
			cnt++
		}
		N >>= 1
	}
	return max
}
