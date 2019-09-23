package Test119

import (
	"fmt"
	"testing"
)

func Test119(t *testing.T) {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	var data = []int{1}
	if rowIndex < 1 {
		return data
	}
	for i := 0; i < rowIndex; i++ {
		data = getData(data)
	}
	return data
}

var data1 = make([]int, 0)

func getData(data []int) []int {
	data1 = []int{}
	data1 = append(data1, 1)
	for i := 1; i < len(data); i++ {
		data1 = append(data1, data[i-1]+data[i])
	}
	data1 = append(data1, 1)
	data = data1
	return data1
}
