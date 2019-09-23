package Test1046

import (
	"sort"
	"testing"
)

func Test1046(t *testing.T) {
	a := []int{2, 7, 4, 1, 8, 1}
	b := lastStoneWeight(a)
	t.Log(b)
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	} else if len(stones) == 2 {
		return abs(stones)
	}
	sort.Ints(stones)
	stones[len(stones)-2] = abs(stones[len(stones)-2:])
	stones = stones[:len(stones)-1]
	return lastStoneWeight(stones)
}
func abs(data []int) int {
	if data[0] > data[1] {
		return data[0] - data[1]
	}
	return data[1] - data[0]
}
