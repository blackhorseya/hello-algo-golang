package chapter_stack_and_queue

import (
	"container/list"
	"fmt"
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestQueue(t *testing.T) {
	/* 初始化队列 */
	// 在 Go 中，将 list 作为队列来使用
	queue := list.New()

	/* 元素入队 */
	queue.PushBack(1)
	queue.PushBack(3)
	queue.PushBack(2)
	queue.PushBack(5)
	queue.PushBack(4)
	fmt.Print("队列 queue = ")
	PrintList(queue)

	/* 访问队首元素 */
	peek := queue.Front()
	fmt.Println("队首元素 peek =", peek.Value)

	/* 元素出队 */
	pop := queue.Front()
	queue.Remove(pop)
	fmt.Print("出队元素 pop = ", pop.Value, "，出队后 queue = ")
	PrintList(queue)

	/* 获取队列的长度 */
	size := queue.Len()
	fmt.Println("队列长度 size =", size)

	/* 判断队列是否为空 */
	isEmpty := queue.Len() == 0
	fmt.Println("队列是否为空 =", isEmpty)
}
