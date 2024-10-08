package pkg

// TreeNode 二元樹節點結構體
type TreeNode struct {
	Val    any       // 節點值
	Height int       // 節點高度
	Left   *TreeNode // 左子樹
	Right  *TreeNode // 右子樹
}

// NewTreeNode 創建一個新的二元樹節點
func NewTreeNode(v any) *TreeNode {
	return &TreeNode{
		Val:    v,
		Height: 0,
		Left:   nil,
		Right:  nil,
	}
}

// 序列化编码规则请参考：
// https://www.hello-algo.com/chapter_tree/array_representation_of_tree/
// 二叉树的数组表示：
// [1, 2, 3, 4, nil, 6, 7, 8, 9, nil, nil, 12, nil, nil, 15]
// 二叉树的链表表示：
//
//	         /——— 15
//	     /——— 7
//	 /——— 3
//	|    \——— 6
//	|        \——— 12
//
// ——— 1
//
//	\——— 2
//	   |    /——— 9
//	    \——— 4
//	        \——— 8

// SliceToTreeDFS 将列表反序列化为二叉树：递归
func SliceToTreeDFS(arr []any, i int) *TreeNode {
	if i < 0 || i >= len(arr) || arr[i] == nil {
		return nil
	}
	root := NewTreeNode(arr[i])
	root.Left = SliceToTreeDFS(arr, 2*i+1)
	root.Right = SliceToTreeDFS(arr, 2*i+2)
	return root
}

// SliceToTree 将切片反序列化为二叉树
func SliceToTree(arr []any) *TreeNode {
	return SliceToTreeDFS(arr, 0)
}

// TreeToSliceDFS 将二叉树序列化为切片：递归
func TreeToSliceDFS(root *TreeNode, i int, res *[]any) {
	if root == nil {
		return
	}
	for i >= len(*res) {
		*res = append(*res, nil)
	}
	(*res)[i] = root.Val
	TreeToSliceDFS(root.Left, 2*i+1, res)
	TreeToSliceDFS(root.Right, 2*i+2, res)
}

// TreeToSlice 将二叉树序列化为切片
func TreeToSlice(root *TreeNode) []any {
	var res []any
	TreeToSliceDFS(root, 0, &res)
	return res
}
