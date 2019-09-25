package Test7

import (
	"testing"
)

func Test(t *testing.T) {

	t.Log(reverse(1534236469))
}

func reverse(x int) int {
	temp := abs(x)
	res := 0
	for temp > 0 {
		res = res*10 + temp%10
		temp = temp / 10
	}
	if res > 2147483647 {
		return 0
	}
	if x < 0 {
		res = -res
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
