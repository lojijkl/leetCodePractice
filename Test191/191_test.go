package Test191

import (
	"testing"
)

func Test191(t *testing.T) {
	var num uint32 = 00000000000000000000000000001011
	aa := hammingWeight(num)
	t.Log(aa)
}

func hammingWeight(num uint32) int {
	cnt = 0
	for num != 0 {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}
	return cnt
}

var cnt int
