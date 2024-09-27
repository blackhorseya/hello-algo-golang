package chapter_tree

import (
	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

type aVLTree struct {
	root *TreeNode
}

func newAVLTree() *aVLTree {
	return &aVLTree{}
}

/* 获取节点高度 */
func (t *aVLTree) height(node *TreeNode) int {
	// 空节点高度为 -1 ，叶节点高度为 0
	if node != nil {
		return node.Height
	}

	return -1
}

/* 更新節點高度 */
func (t *aVLTree) updateHeight(node *TreeNode) {
	leftHeight := t.height(node.Left)
	rightHeight := t.height(node.Right)

	// 節點高度等於最高子樹高度 + 1
	if leftHeight > rightHeight {
		node.Height = leftHeight + 1
	} else {
		node.Height = rightHeight + 1
	}
}

/* 獲取平衡因子 */
func (t *aVLTree) balanceFactor(node *TreeNode) int {
	// 空節點平衡因子為 0
	if node == nil {
		return 0
	}
	// 節點平衡因子 = 左子樹高度 - 右子樹高度
	return t.height(node.Left) - t.height(node.Right)
}

/* 右旋操作 */
func (t *aVLTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right

	// 以 child 為原點，將 node 向右旋轉
	child.Right = node
	node.Left = grandChild

	// 更新節點高度
	t.updateHeight(node)
	t.updateHeight(child)

	return child
}

/* 左旋操作 */
func (t *aVLTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left

	// 以 child 為原點，將 node 向左旋轉
	child.Left = node
	node.Right = grandChild

	// 更新節點高度
	t.updateHeight(node)
	t.updateHeight(child)

	return child
}

/* 執行旋轉操作，使該子樹重新恢復平衡 */
func (t *aVLTree) rotate(node *TreeNode) *TreeNode {
	// 獲取節點 node 的平衡因子
	// Go 推薦短變數，這裡 bf 指代 t.balanceFactor
	bf := t.balanceFactor(node)

	// 左偏樹
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			return t.rightRotate(node)
		} else {
			node.Left = t.leftRotate(node.Left)
			return t.rightRotate(node)
		}
	}

	// 右偏樹
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			return t.leftRotate(node)
		} else {
			node.Right = t.rightRotate(node.Right)
			return t.leftRotate(node)
		}
	}

	return node
}

/* 插入節點 */
func (t *aVLTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

/* 遞迴插入節點（輔助函式） */
func (t *aVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}

	// 1. 查詢插入位置並插入節點
	if val < node.Val.(int) {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val.(int) {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		// 重複節點，不處理
		return node
	}

	// 更新節點高度
	t.updateHeight(node)

	// 2. 旋轉操作，使該子樹重新恢復平衡
	return t.rotate(node)
}

/* 刪除節點 */
func (t *aVLTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

/* 遞迴刪除節點（輔助函式） */
func (t *aVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	// 1. 查詢刪除位置並刪除節點
	if val < node.Val.(int) {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val.(int) {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				// 子節點數量 = 0 ，直接刪除 node 並返回
				return nil
			} else {
				// 子節點數量 = 1 ，直接刪除 node
				node = child
			}
		} else {
			// 子節點數量 = 2 ，則將中序走訪的下個節點刪除，並用該節點替換當前節點
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val.(int))
			node.Val = temp.Val
		}
	}

	// 更新節點高度
	t.updateHeight(node)

	// 2. 旋轉操作，使該子樹重新恢復平衡
	return t.rotate(node)
}

/* 查找节点 */
func (t *aVLTree) search(val int) *TreeNode {
	cur := t.root
	// 循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.Val.(int) < val {
			// 目标节点在 cur 的右子树中
			cur = cur.Right
		} else if cur.Val.(int) > val {
			// 目标节点在 cur 的左子树中
			cur = cur.Left
		} else {
			// 找到目标节点，跳出循环
			break
		}
	}
	// 返回目标节点
	return cur
}
