package Test922

import "testing"

func Test(t *testing.T) {
	a := []int{4, 2, 5, 7}
	//a:=[]int{3,4}
	b := sortArrayByParityII2(a)
	t.Log(b)
}

func sortArrayByParityII(A []int) []int {
	i, j := 0, 1
	for i < len(A) && j < len(A) {
		for i < len(A) && A[i]%2 == 0 {
			i = i + 2
		}
		for j < len(A) && A[j]%2 == 1 {
			j = j + 2
		}
		if i < len(A) && j < len(A) {
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}

func sortArrayByParityII2(A []int) []int {
	res := make([]int, len(A))
	var index1, index2 int
	index2 = 1
	for _, v := range A {
		if v%2 == 0 {
			res[index1] = v
			index1 = index1 + 2
		} else {
			res[index2] = v
			index2 = index2 + 2
		}
	}
	return res
}
