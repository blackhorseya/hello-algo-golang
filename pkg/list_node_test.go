package pkg

import (
	"testing"
)

func TestListNode(t *testing.T) {
	arr := []int{2, 3, 5, 6, 7}
	head := ArrayToLinkedList(arr)

	PrintLinkedList(head)
}
