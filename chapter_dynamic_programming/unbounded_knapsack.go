package chapter_dynamic_programming

/* 完全背包：動態規劃 */
func unboundedKnapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)
	// 初始化 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, cap+1)
	}
	// 狀態轉移
	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			if wgt[i-1] > c {
				// 若超過背包容量，則不選物品 i
				dp[i][c] = dp[i-1][c]
			} else {
				// 不選和選物品 i 這兩種方案的較大值
				dp[i][c] = max(dp[i-1][c], dp[i][c-wgt[i-1]]+val[i-1])
			}
		}
	}
	return dp[n][cap]
}

/* 完全背包：空間最佳化後的動態規劃 */
func unboundedKnapsackDPComp(wgt, val []int, cap int) int {
	n := len(wgt)
	// 初始化 dp 表
	dp := make([]int, cap+1)
	// 狀態轉移
	for i := 1; i <= n; i++ {
		for c := wgt[i-1]; c <= cap; c++ {
			dp[c] = max(dp[c], dp[c-wgt[i-1]]+val[i-1])
		}
	}
	return dp[cap]
}
