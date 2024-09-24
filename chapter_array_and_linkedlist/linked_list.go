package chapter_array_and_linkedlist

import (
	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func insertNode(n0 *ListNode, P *ListNode) {
	n1 := n0.Next
	P.Next = n1
	n0.Next = P
}

func removeItem(n0 *ListNode) {
	if n0.Next == nil {
		return
	}

	P := n0.Next
	n1 := P.Next
	n0.Next = n1
}

func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head.Next = head
	}

	return head
}

func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}

	return -1
}
