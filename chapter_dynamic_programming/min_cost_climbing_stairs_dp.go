package chapter_dynamic_programming

/* 爬樓梯最小代價：動態規劃 */
func minCostClimbingStairsDP(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return cost[n]
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 初始化 dp 表，用於儲存子問題的解
	dp := make([]int, n+1)
	// 初始狀態：預設最小子問題的解
	dp[1] = cost[1]
	dp[2] = cost[2]
	// 狀態轉移：從較小子問題逐步求解較大子問題
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return dp[n]
}

/* 爬樓梯最小代價：空間最佳化後的動態規劃 */
func minCostClimbingStairsDPComp(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return cost[n]
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 初始狀態：預設最小子問題的解
	a, b := cost[1], cost[2]
	// 狀態轉移：從較小子問題逐步求解較大子問題
	for i := 3; i <= n; i++ {
		a, b = b, min(a, b)+cost[i]
	}

	return b
}
