package chapter_dynamic_programming

/* 爬樓梯：記憶化搜尋 */
func climbingStairsDFSMem(n int) int {
	// mem[i] 記錄爬到第 i 階的方案總數，-1 代表無記錄
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = -1
	}

	return dfsMem(n, mem)
}

/* 記憶化搜尋 */
func dfsMem(i int, mem []int) int {
	// 已知 dp[1] 和 dp[2] ，返回之
	if i == 1 || i == 2 {
		return i
	}

	// 若 mem[i] 有記錄，直接返回
	if mem[i] != -1 {
		return mem[i]
	}

	// dp[i] = dp[i-1] + dp[i-2]
	count := dfsMem(i-1, mem) + dfsMem(i-2, mem)
	// 更新 mem[i]
	mem[i] = count
	return count
}
