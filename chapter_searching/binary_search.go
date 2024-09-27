package chapter_searching

/* 二分搜尋（雙閉區間） */
func binarySearch(nums []int, target int) int {
	// 定義左右邊界
	left, right := 0, len(nums)-1

	// 當左邊界小於等於右邊界時, 代表還有元素可以檢查
	for left <= right {
		// 取得中間索引與值
		mid := left + (right-left)/2
		val := nums[mid]

		if val < target { // val < target 代表目標在右半邊 (mid 已經檢查過, 所以 left = mid + 1)
			left = mid + 1
		} else if val > target { // val > target 代表目標在左半邊 (mid 已經檢查過, 所以 right = mid - 1)
			right = mid - 1
		} else { // val == target 代表找到目標, 回傳索引
			return mid
		}
	}

	// 找不到, 回傳 -1
	return -1
}
