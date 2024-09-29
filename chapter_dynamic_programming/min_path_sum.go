package chapter_dynamic_programming

import (
	"math"
)

/* 最小路徑和：暴力搜尋 */
func minPathSumDFS(grid [][]int, i, j int) int {
	// 若為左上角單元格，則終止搜尋
	if i == 0 && j == 0 {
		return grid[0][0]
	}

	// 若行列索引越界，則返回 +∞ 代價
	if i < 0 || j < 0 {
		return math.MaxInt
	}

	// 計算從左上角到 (i-1, j) 和 (i, j-1) 的最小路徑代價
	up := minPathSumDFS(grid, i-1, j)
	left := minPathSumDFS(grid, i, j-1)

	// 返回從左上角到 (i, j) 的最小路徑代價
	return int(math.Min(float64(up), float64(left))) + grid[i][j]
}

/* 最小路徑和：記憶化搜尋 */
func minPathSumDFSMem(grid, mem [][]int, i, j int) int {
	// 若為左上角單元格，則終止搜尋
	if i == 0 && j == 0 {
		return grid[0][0]
	}

	// 若行列索引越界，則返回 +∞ 代價
	if i < 0 || j < 0 {
		return math.MaxInt
	}

	// 若已經計算過，則直接返回
	if mem[i][j] != -1 {
		return mem[i][j]
	}

	// 計算從左上角到 (i-1, j) 和 (i, j-1) 的最小路徑代價
	up := minPathSumDFSMem(grid, mem, i-1, j)
	left := minPathSumDFSMem(grid, mem, i, j-1)
	// 記錄並返回左上角到 (i, j) 的最小路徑代價
	mem[i][j] = int(math.Min(float64(left), float64(up))) + grid[i][j]
	return mem[i][j]
}
