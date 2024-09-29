package pkg

// ListNode represents a node in a singly linked list
type ListNode struct {
	Next *ListNode
	Val  int
}

// NewListNode creates a new ListNode
func NewListNode(val int) *ListNode {
	return &ListNode{
		Next: nil,
		Val:  val,
	}
}

// ArrayToLinkedList 将数组反序列化为链表
func ArrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := NewListNode(arr[0])
	cur := head
	for _, val := range arr {
		cur.Next = NewListNode(val)
		cur = cur.Next
	}

	return head.Next
}
