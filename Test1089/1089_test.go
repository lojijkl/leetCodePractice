package Test1089

import "testing"

func Test(t *testing.T) {
	//a:=[]int{1,0,2,3,0,4,5,0}
	a := []int{1, 2, 3}
	duplicateZeros(a)
	t.Log(a)
}
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
			i++
		}
	}
}
