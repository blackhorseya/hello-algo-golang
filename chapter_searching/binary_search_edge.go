package chapter_searching

/* 二分搜尋最左一個 target */
func binarySearchLeftEdge(nums []int, target int) int {
	// 等價於查詢 target 的插入點
	i := binarySearchInsertion(nums, target)

	// 如果 i 等於 nums 的長度，或者 nums[i] 不等於 target，則表示 nums 中不存在 target
	if i == len(nums) || nums[i] != target {
		return -1
	}

	// 這裡的 i 就是最左一個 target 的索引
	return i
}

/* 二分查找最右一个 target */
func binarySearchRightEdge(nums []int, target int) int {
	// 转化为查找最左一个 target + 1
	i := binarySearchInsertion(nums, target+1)
	// j 指向最右一个 target ，i 指向首个大于 target 的元素
	j := i - 1
	// 未找到 target ，返回 -1
	if j == -1 || nums[j] != target {
		return -1
	}
	// 找到 target ，返回索引 j
	return j
}
