package chapter_tree

import (
	"fmt"
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestLevelOrder(t *testing.T) {
	/* 初始化二叉树 */
	// 这里借助了一个从数组直接生成二叉树的函数
	root := SliceToTree([]any{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\n初始化二叉树: ")
	PrintTree(root)

	// 层序遍历
	nums := levelOrder(root)
	fmt.Println("\n层序遍历的节点打印序列 =", nums)
}
