package chapter_divide_and_conquer

import (
	"fmt"
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 2, 1, 7}
	inorder := []int{9, 3, 1, 2, 7}
	fmt.Print("前序遍历 = ")
	PrintSlice(preorder)
	fmt.Print("中序遍历 = ")
	PrintSlice(inorder)

	root := buildTree(preorder, inorder)
	fmt.Println("构建的二叉树为：")
	PrintTree(root)
}
