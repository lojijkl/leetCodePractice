package Test125

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	a := "9"
	fmt.Println([]rune(a))
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	t.Log(isPalindrome("\"0P\""))
}

func isPalindrome(s string) bool {
	aa := []rune(s)
	j := len(aa) - 1
	for i := 0; i < j; i, j = i+1, j-1 {
		for !checkUseful(rune(aa[i])) && i < j {
			i++
		}
		for !checkUseful(rune(aa[j])) && i < j {
			j--
		}
		if aa[i] != aa[j] && aa[i] != aa[j]^32 {
			return false
		}
	}
	return true
}

func checkUseful(s rune) bool {
	if (s >= 65 && s <= 90) || (s >= 97 && s <= 122) || (s >= 48 && s <= 57) {
		return true
	}
	return false
}
