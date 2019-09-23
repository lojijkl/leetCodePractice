package Test2

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
	}
	l2 := &ListNode{
		Val: 9,
	}
	data := addTwoNumbers(l1, l2)
	for data != nil {
		fmt.Println(data.Val)
		data = data.Next
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		if swit {
			return &ListNode{Val: 1}
		} else {
			return nil
		}
	}
	if l1 == nil {
		l1 = &ListNode{Val: 0}
	}
	if l2 == nil {
		l2 = &ListNode{Val: 0}
	}
	val := l1.Val + l2.Val
	if swit {
		val += 1
		swit = false
	}
	node := &ListNode{}

	if val > 9 {
		swit = true
		node.Val = val % 10
	} else {
		node.Val = val
	}
	node.Next = addTwoNumbers(l1.Next, l2.Next)
	return node
}

var swit = false
