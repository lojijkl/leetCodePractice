package Test908

import (
	"sort"
	"testing"
)

func Test908(t *testing.T) {
	//A := []int{0,10}
	//K := 2
	//t.Log(smallestRangeI(A,K))
	//a:=[]int{2,7,10,9,8,9}
	//fmt.Println(a[0:3])
	//a:=[]int{1,2,3}
	a := []int{7, 4, 8, 9, 7, 7, 5}
	t.Log(movesToMakeZigzag(a))
	t.Log(a)
}

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	if A[0]+K > A[len(A)-1]-K {
		return 0
	}
	return A[len(A)-1] - A[0] - 2*K
}

func movesToMakeZigzag(nums []int) int {
	var cnt int
	l := len(nums)
	var up, down int
	for i := 1; i < l; i = i + 2 {
		if i+2 <= l {
			if nums[i-1] > nums[i] && nums[i] < nums[i+1] {
				down++
			} else if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
				up++
			}
		}
	}
	for i := 1; i < l; i = i + 2 {
		if i+2 <= l {
			cnt += moves(nums[i-1 : i+2])
		}
	}
	return cnt
}
func moves(data []int) int {
	if len(data) < 3 {
		return 0
	}
	if (data[0] >= data[1] && data[1] >= data[2]) || (data[0] <= data[1] && data[1] <= data[2]) {
		a := data[0] - data[1]
		b := data[1] - data[2]
		var cnt int
		if abs(a) > abs(b) {
			cnt = abs(data[2]-data[1]) + 1
			data[2] = data[2] - cnt
		} else {
			cnt = abs(data[1]-data[0]) + 1
			data[0] = data[0] - cnt
		}
		return cnt
	}
	return 0
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
