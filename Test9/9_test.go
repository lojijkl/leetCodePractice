package Test9

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Log(isPalindrome(121))
	t.Log(isPalindrome(-121))
	t.Log(isPalindrome(10))
}

//转字符串
func isPalindrome1(x int) bool {
	aa := fmt.Sprintf("%v", x)
	l := len(aa)
	for i := 0; i < l+1/2; i++ {
		if aa[i] != aa[l-i-1] {
			return false
		}
	}

	return true
}

//进阶
func isPalindrome(x int) bool {
	temp := x
	reverse := 0
	if x < 0 {
		return false
	}
	for temp > 0 {
		reverse = reverse*10 + temp%10
		temp /= 10
	}

	for reverse > 0 {
		if reverse%10 != x%10 {
			return false
		}
		reverse /= 10
		x /= 10
	}
	return true
}
