package Test784

import (
	"testing"
)

func Test(t *testing.T) {
	s := "a1b2"
	res := letterCasePermutation2(s)
	t.Log(res)
}

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return nil
	}
	res := []string{}
	if S[0] >= 65 {
		res = append(res, string(S[0]^(1<<5)))
	}
	res = append(res, string(S[0]))
	for j := 1; j < len(S); j++ {
		l := len(res)
		for i := 0; i < l; i++ {
			if S[j] >= 65 {
				res = append(res, res[i]+string(S[j]^(1<<5)))
			}
			res[i] = res[i] + string(S[j])
		}
	}
	return res
}
