package chapter_searching

/* 二分搜尋插入點（無重複元素） */
func binarySearchInsertionSimple(nums []int, target int) int {
	// 定義左右邊界 [0, n-1]
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

	return left
}

/* 二分搜尋插入點（存在重複元素） */
func binarySearchInsertion(nums []int, target int) int {
	// 初始化雙閉區間 [0, n-1]
	i, j := 0, len(nums)-1

	for i <= j {
		// 計算中點索引 m
		m := i + (j-i)/2
		if nums[m] < target {
			// target 在區間 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target {
			// target 在區間 [i, m-1] 中
			j = m - 1
		} else {
			// 首個小於 target 的元素在區間 [i, m-1] 中
			j = m - 1
		}
	}

	// 返回插入點 i
	return i
}
