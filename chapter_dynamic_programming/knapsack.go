package chapter_dynamic_programming

/* 0-1 背包：暴力搜尋 */
func knapsackDFS(wgt, val []int, i, c int) int {
	// 若已選完所有物品或背包無剩餘容量，則返回價值 0
	if i == 0 || c == 0 {
		return 0
	}

	// 若超過背包容量，則只能選擇不放入背包
	if wgt[i-1] > c {
		return knapsackDFS(wgt, val, i-1, c)
	}

	// 計算不放入和放入物品 i 的最大價值
	no := knapsackDFS(wgt, val, i-1, c)
	yes := knapsackDFS(wgt, val, i-1, c-wgt[i-1]) + val[i-1]

	return max(no, yes)
}

/* 0-1 背包：記憶化搜尋 */
func knapsackDFSMem(wgt, val []int, mem [][]int, i, c int) int {
	// 若已選完所有物品或背包無剩餘容量，則返回價值 0
	if i == 0 || c == 0 {
		return 0
	}

	// 若已計算過，則直接返回結果
	if mem[i][c] != -1 {
		return mem[i][c]
	}

	// 若超過背包容量，則只能選擇不放入背包
	if wgt[i-1] > c {
		return knapsackDFSMem(wgt, val, mem, i-1, c)
	}

	// 計算不放入和放入物品 i 的最大價值
	no := knapsackDFSMem(wgt, val, mem, i-1, c)
	yes := knapsackDFSMem(wgt, val, mem, i-1, c-wgt[i-1]) + val[i-1]

	mem[i][c] = max(no, yes)
	return mem[i][c]
}
