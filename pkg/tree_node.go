package pkg

// TreeNode 二元樹節點結構體
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 創建一個新的二元樹節點
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Left:  nil, // 左子節點指標
		Right: nil, // 右子節點指標
		Val:   v,   // 節點值
	}
}
