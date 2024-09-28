package chapter_backtracking

import (
	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

/* 前序走訪：例題一 */
func preOrderI(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	if (root.Val).(int) == 7 {
		// 記錄解
		*res = append(*res, root)
	}
	preOrderI(root.Left, res)
	preOrderI(root.Right, res)
}
