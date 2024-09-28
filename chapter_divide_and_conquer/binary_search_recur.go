package chapter_divide_and_conquer

func dfs(nums []int, target, i, j int) int {
	if i > j {
		return -1
	}
	mid := i + (j-i)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return dfs(nums, target, i, mid-1)
	} else {
		return dfs(nums, target, mid+1, j)
	}
}

func binarySearch(nums []int, target int) int {
	return dfs(nums, target, 0, len(nums)-1)
}
