package Test121

import (
	"testing"
)

func Test(t *testing.T) {
	//t.Log(maxProfit([]int{7,1,5,3,6,4}))
	t.Log(maxProfit([]int{}))
	//t.Log(maxProfit([]int{2,4,1}))
	//t.Log(maxProfit([]int{3,3,5,0,0,3,1,4}))
}

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
