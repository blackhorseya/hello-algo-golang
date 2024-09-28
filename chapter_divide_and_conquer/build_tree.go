package chapter_divide_and_conquer

import (
	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

/* 構建二元樹：分治 */
func dfsBuildTree(preorder []int, inorderMap map[int]int, index, left, right int) *TreeNode {
	// 子樹區間為空時終止
	if right-left < 0 {
		return nil
	}

	// 建立根節點
	root := NewTreeNode(preorder[index])
	// 獲取根節點在 inorder 中的索引
	inorderIndex := inorderMap[preorder[index]]
	// 遞歸構建左子樹
	root.Left = dfsBuildTree(preorder, inorderMap, index+1, left, inorderIndex-1)
	// 遞歸構建右子樹
	root.Right = dfsBuildTree(preorder, inorderMap, index+inorderIndex-left+1, inorderIndex+1, right)

	return root
}

func buildTree(preorder, inorder []int) *TreeNode {
	// 初始化雜湊表，儲存 inorder 元素到索引的對映
	inorderMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderMap[v] = i
	}

	return dfsBuildTree(preorder, inorderMap, 0, 0, len(inorder)-1)
}
