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
