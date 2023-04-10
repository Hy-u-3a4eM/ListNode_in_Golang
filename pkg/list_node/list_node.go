package list_node

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{Next: nil, Val: val}
}

func Push(next *ListNode, val int) *ListNode {
	return &ListNode{Next: next, Val: val}
}

func SliceToListNode(slice []int) *ListNode {
	current := New(slice[0])
	for i := 1; i < len(slice); i += 1 {
		current = Push(current, slice[i])
	}

	return current
}

func PrintAll(node *ListNode) {
	fmt.Print("[")
	for {
		if node == nil {
			fmt.Print("]\n")
			return
		}

		if node.Next == nil {
			fmt.Print(node.Val)
		} else {
			fmt.Print(node.Val, " ")
		}

		node = node.Next
	}
}
