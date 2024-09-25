package chapter_stack_and_queue

import (
	"container/list"
	"fmt"
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestDeque(t *testing.T) {
	/* 初始化双端队列 */
	deque := list.New()

	/* 元素入队 */
	deque.PushBack(2)
	deque.PushBack(5)
	deque.PushBack(4)
	deque.PushFront(3)
	deque.PushFront(1)
	fmt.Print("双向队列 deque = ")
	PrintList(deque)

	/* 访问元素 */
	front := deque.Front()
	fmt.Println("队首元素 front =", front.Value)
	rear := deque.Back()
	fmt.Println("队尾元素 rear =", rear.Value)

	/* 元素出队 */
	deque.Remove(front)
	fmt.Print("队首出队元素 front = ", front.Value, "，队首出队后 deque = ")
	PrintList(deque)
	deque.Remove(rear)
	fmt.Print("队尾出队元素 rear = ", rear.Value, "，队尾出队后 deque = ")
	PrintList(deque)

	/* 获取双向队列的长度 */
	size := deque.Len()
	fmt.Println("双向队列长度 size =", size)

	/* 判断双向队列是否为空 */
	isEmpty := deque.Len() == 0
	fmt.Println("双向队列是否为空 =", isEmpty)
}
