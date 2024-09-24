package chapter_stack_and_queue

import (
	"fmt"
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestStack(t *testing.T) {
	/* 初始化栈 */
	// 在 Go 中，推荐将 Slice 当作栈来使用
	var stack []int

	/* 元素入栈 push */
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	stack = append(stack, 5)
	stack = append(stack, 4)
	fmt.Print("栈 stack = ")
	PrintSlice(stack)

	/* 访问栈顶元素 peak */
	peek := stack[len(stack)-1]
	fmt.Println("栈顶元素 peek =", peek)

	/* 元素出栈 pop */
	pop := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Print("出栈元素 pop = ", pop, "，出栈后 stack = ")
	PrintSlice(stack)

	/* 获取栈的长度 */
	size := len(stack)
	fmt.Println("栈的长度 size =", size)

	/* 判断是否为空 */
	isEmpty := len(stack) == 0
	fmt.Println("栈是否为空 =", isEmpty)
}
