package main

import (
	"fmt"

	"ListNode/pkg/list_node"
)

func addTwoNumbers(l1 *list_node.ListNode, l2 *list_node.ListNode) *list_node.ListNode {
	var result *list_node.ListNode = nil
	current1, current2 := l1, l2
	for {
		result = list_node.Push(result, current1.Val+current2.Val)
		current1, current2 = current1.Next, current2.Next
	}

	return result
}

func main() {
	fmt.Println(
		addTwoNumbers(list_node.SliceToListNode([]int{2, 4, 3}),
			list_node.SliceToListNode([]int{5, 6, 4})))
}
