package Test561

import (
	"fmt"
	"testing"
)

func Test561(t *testing.T) {
	s := []int{9, 4, 6, 7, 6, 5}
	quickSort1Desc(s)
	//a:=arrayPairSum(s)
	t.Log(s)
}

func arrayPairSum(nums []int) int {
	quickSort(nums)
	sums := 0
	for i := 0; i < len(nums); i = i + 2 {
		sums += nums[i]
	}
	return sums
}
func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	tempValue, pivot := data[0], 1
	head, tail := 0, len(data)-1

	for head < tail {
		if tempValue < data[pivot] {
			data[pivot], data[tail] = data[tail], data[pivot]
			tail--
		} else {
			data[pivot], data[head] = data[head], data[pivot]
			head++
			pivot++
		}
	}
	data[head] = tempValue
	quickSort(data[:head])
	quickSort(data[head+1:])
}

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort1(values)
}

func quickSort1(data []int) {
	if len(data) <= 1 {
		return
	}
	key := data[0]
	left, right := 0, len(data)-1
	for left < right {
		fmt.Println(left, right)
		for left < right && key <= data[right] {
			right--
		}
		data[left] = data[right]
		for left < right && key >= data[left] {
			left++
		}
		data[right] = data[left]
	}
	data[left] = key
	quickSort1(data[:left])
	quickSort1(data[left+1:])
}

func quickSort1Desc(data []int) {
	fmt.Println(data)
	if len(data) <= 1 {
		return
	}
	key := data[0]
	left, right := 0, len(data)-1
	for left < right {
		for left < right && key >= data[right] {
			right--
		}
		data[left] = data[right]
		for left < right && key <= data[left] {
			left++
		}
		data[right] = data[left]
	}
	data[left] = key
	quickSort1Desc(data[:left])
	quickSort1Desc(data[left+1:])
}
