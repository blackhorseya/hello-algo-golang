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

/* 二分搜尋（左閉右開區間） */
func binarySearchLCRO(nums []int, target int) int {
	// 初始化左閉右開區間 [0, n) ，即 i, j 分別指向陣列首元素、尾元素+1
	i, j := 0, len(nums)
	// 迴圈，當搜尋區間為空時跳出（當 i = j 時為空）
	for i < j {
		m := i + (j-i)/2      // 計算中點索引 m
		if nums[m] < target { // 此情況說明 target 在區間 [m+1, j) 中
			i = m + 1
		} else if nums[m] > target { // 此情況說明 target 在區間 [i, m) 中
			j = m
		} else { // 找到目標元素，返回其索引
			return m
		}
	}
	// 未找到目標元素，返回 -1
	return -1
}
