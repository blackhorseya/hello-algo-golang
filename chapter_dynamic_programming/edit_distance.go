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
