package chapter_dynamic_programming

/* 零錢兌換：動態規劃 */
func coinChangeDP(coins []int, amt int) int {
	n := len(coins)

	// 初始化 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amt+1)
	}

	// 狀態轉移：首行首列
	for a := 1; a <= amt; a++ {
		dp[0][a] = amt + 1 // 這裡的 amt + 1 可以理解為無限大
	}

	// 狀態轉移：填表
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				dp[i][a] = dp[i-1][a]
			} else {
				dp[i][a] = min(dp[i-1][a], dp[i][a-coins[i-1]]+1)
			}
		}
	}

	res := dp[n][amt]
	if res > amt {
		res = -1
	}

	return res
}

/* 零錢兌換：動態規劃 */
func coinChangeDPComp(coins []int, amt int) int {
	n := len(coins)

	// 初始化 dp 表
	dp := make([]int, amt+1)

	// 狀態轉移：首行
	for a := 1; a <= amt; a++ {
		dp[a] = amt + 1 // 這裡的 amt + 1 可以理解為無限大
	}

	// 狀態轉移：填表
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] <= a {
				dp[a] = min(dp[a], dp[a-coins[i-1]]+1)
			}
		}
	}

	res := dp[amt]
	if res > amt {
		res = -1
	}

	return res
}
