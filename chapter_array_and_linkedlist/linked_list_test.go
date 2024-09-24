package chapter_array_and_linkedlist

import (
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestLinkedList(t *testing.T) {
	/* 初始化链表 1 -> 3 -> 2 -> 5 -> 4 */
	// 初始化各个节点
	n0 := NewListNode(1)
	n1 := NewListNode(3)
	n2 := NewListNode(2)
	n3 := NewListNode(5)
	n4 := NewListNode(4)

	// 构建节点之间的引用
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	// fmt.Println("初始化的链表为")
	// PrintLinkedList(n0)
}
