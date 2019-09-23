package Test171

import "testing"

func Test(t *testing.T) {
	t.Log(titleToNumber("Z"))
}

func titleToNumber(s string) int {
	res := 0
	for _, v := range s {
		res = res*26 + int(v^64)
	}
	return res
}
