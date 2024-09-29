package chapter_dynamic_programming

/* 编辑距离：暴力搜索 */
func editDistanceDFS(s string, t string, i int, j int) int {
	// 若 s 和 t 都为空，则返回 0
	if i == 0 && j == 0 {
		return 0
	}

	// 若 s 为空，t 不为空，则返回 t 的长度
	if i == 0 {
		return j
	}

	// 若 t 为空，s 不为空，则返回 s 的长度
	if j == 0 {
		return i
	}

	// 若 s 和 t 的最后一个字符相同，则不需要编辑
	if s[i-1] == t[j-1] {
		return editDistanceDFS(s, t, i-1, j-1)
	}

	// 若 s 和 t 的最后一个字符不同，则有三种操作
	// 1. 插入操作：在 s 的末尾插入 t 的最后一个字符
	// 2. 删除操作：删除 s 的最后一个字符
	// 3. 替换操作：将 s 的最后一个字符替换为 t 的最后一个字符
	return min(
		editDistanceDFS(s, t, i, j-1)+1,
		editDistanceDFS(s, t, i-1, j)+1,
		editDistanceDFS(s, t, i-1, j-1)+1,
	)
}

/* 编辑距离：记忆化搜索 */
func editDistanceDFSMem(s string, t string, mem [][]int, i int, j int) int {
	// 若 s 和 t 都为空，则返回 0
	if i == 0 && j == 0 {
		return 0
	}

	// 若 s 为空，t 不为空，则返回 t 的长度
	if i == 0 {
		return j
	}

	// 若 t 为空，s 不为空，则返回 s 的长度
	if j == 0 {
		return i
	}

	// 若 s 和 t 的最后一个字符相同，则不需要编辑
	if s[i-1] == t[j-1] {
		return editDistanceDFSMem(s, t, mem, i-1, j-1)
	}

	// 若 s 和 t 的最后一个字符不同，则有三种操作
	// 1. 插入操作：在 s 的末尾插入 t 的最后一个字符
	// 2. 删除操作：删除 s 的最后一个字符
	// 3. 替换操作：将 s 的最后一个字符替换为 t 的最后一个字符
	if mem[i][j] != -1 {
		return mem[i][j]
	}

	mem[i][j] = min(
		editDistanceDFSMem(s, t, mem, i, j-1)+1,
		editDistanceDFSMem(s, t, mem, i-1, j)+1,
		editDistanceDFSMem(s, t, mem, i-1, j-1)+1,
	)
	return mem[i][j]
}

/* 编辑距离：动态规划 */
func editDistanceDP(s string, t string) int {
	n := len(s)
	m := len(t)

	// 初始化二维数组
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 边界条件
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	// 动态规划
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 若 s 和 t 的最后一个字符相同，则不需要编辑
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 若 s 和 t 的最后一个字符不同，则有三种操作
				// 1. 插入操作：在 s 的末尾插入 t 的最后一个字符
				// 2. 删除操作：删除 s 的最后一个字符
				// 3. 替换操作：将 s 的最后一个字符替换为 t 的最后一个字符
				dp[i][j] = min(
					dp[i][j-1]+1,
					dp[i-1][j]+1,
					dp[i-1][j-1]+1,
				)
			}
		}
	}

	return dp[n][m]
}

func editDistanceDPComp(s string, t string) int {
	n := len(s)
	m := len(t)

	// 优化空间复杂度
	dp := make([]int, m+1)

	// 边界条件
	for j := 0; j <= m; j++ {
		dp[j] = j
	}

	// 动态规划/状态转移方程
	for i := 1; i <= n; i++ {
		leftUp := dp[0] // 暫存 dp[i-1][j-1] 的值
		dp[0] = i
		for j := 1; j <= m; j++ {
			// 保存 dp[i-1][j] 的值
			temp := dp[j]
			// 若 s 和 t 的最后一个字符相同，则不需要编辑
			if s[i-1] == t[j-1] {
				dp[j] = leftUp
			} else {
				// 若 s 和 t 的最后一个字符不同，则有三种操作
				// 1. 插入操作：在 s 的末尾插入 t 的最后一个字符
				// 2. 删除操作：删除 s 的最后一个字符
				// 3. 替换操作：将 s 的最后一个字符替换为 t 的最后一个字符
				dp[j] = min(
					dp[j-1]+1,
					dp[j]+1,
					leftUp+1,
				)
			}
			leftUp = temp
		}
	}

	return dp[m]
}
