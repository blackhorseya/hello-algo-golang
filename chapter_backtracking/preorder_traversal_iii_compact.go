package chapter_backtracking

import (
	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

/* 前序走訪：例題三 */
func preOrderIII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	// 剪枝
	if root == nil || root.Val.(int) == 3 {
		return
	}
	// 嘗試
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		// 記錄解
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderIII(root.Left, res, path)
	preOrderIII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}
