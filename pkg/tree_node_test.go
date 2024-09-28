package pkg

import (
	"fmt"
	"testing"
)

func TestTreeNode(t *testing.T) {
	arr := []any{1, 2, 3, nil, 5, 6, nil}
	node := SliceToTree(arr)

	// print tree
	PrintTree(node)

	// tree to arr
	fmt.Println(TreeToSlice(node))
}
