package Test443

import "testing"

func Test(t *testing.T) {
	b := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	compress(b)
}
func compress(chars []byte) int {
	var temp byte
	var cnt int
	for i := 0; i < len(chars); i++ {
		if chars[i] == temp {
			cnt++
		} else {
			temp = chars[i]
			chars = chars[:i-cnt]
		}
	}
}
