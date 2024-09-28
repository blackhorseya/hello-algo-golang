package chapter_backtracking

import (
	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

/* 前序走訪：例題二 */
func preOrderII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil {
		return
	}

	// 嘗試
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		// 記錄解
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)

	// 回溯/回退
	*path = (*path)[:len(*path)-1]
}
