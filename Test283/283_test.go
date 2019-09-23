package Test283

import "testing"

func Test(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	//a:=[]int{0,1}
	moveZeroes(a)
	t.Log(a)
}
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := 0
	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
