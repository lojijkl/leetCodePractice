package Test169

import "testing"

func Test(t *testing.T) {
	a := []int{3, 2, 3}
	res := majorityElement(a)
	t.Log(res)
}

/*func majorityElement(nums []int) int {
	maps:=map[int]int{}
	l:=len(nums)
	for _,v:=range nums{
		maps[v]++
		if maps[v]<<1>l{
			return v
		}
	}
	return 0
}*/

func majorityElement(nums []int) int {
	max := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if max != nums[i] {
			cnt--
			if cnt == 0 {
				max = nums[i+1]
			}
		} else {
			cnt++
		}
	}
	return max
}
