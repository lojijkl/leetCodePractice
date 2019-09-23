package Test961

import "testing"

func Test961(t *testing.T) {
	a := []int{1, 2, 3, 3}
	data := repeatedNTimes(a)
	t.Log(data)
}

func repeatedNTimes(A []int) int {
	maps = map[int]bool{}
	for _, v := range A {
		if maps[v] {
			return v
		} else {
			maps[v] = true
		}
	}
	return 0
}

var maps = map[int]bool{}
