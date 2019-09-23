package Test884

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	A := "this apple is sweet"
	B := "this apple is sour"
	res := uncommonFromSentences(A, B)
	t.Log(res)
}

func uncommonFromSentences(A string, B string) []string {
	a := strings.Split(A, " ")
	b := strings.Split(B, " ")
	mapA := map[string]int{}
	for _, v := range a {
		mapA[v] = mapA[v] + 1
	}
	for _, v := range b {
		mapA[v] = mapA[v] + 1
	}
	res := []string{}
	for k, v := range mapA {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
