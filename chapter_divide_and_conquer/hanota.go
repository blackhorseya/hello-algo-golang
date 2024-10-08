package chapter_divide_and_conquer

import (
	"container/list"
	"fmt"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

/* 移動一個圓盤 */
func move(src, tar *list.List) {
	// 從 src 頂部拿出一個圓盤
	pan := src.Back()
	// 將圓盤放入 tar 頂部
	tar.PushBack(pan.Value)
	// 移除 src 頂部圓盤
	src.Remove(pan)
}

/* 求解河內塔問題 f(i) */
func dfsHanota(i int, src, buf, tar *list.List) {
	// 若 src 只剩下一個圓盤，則直接將其移到 tar
	if i == 1 {
		move(src, tar)
		return
	}
	// 子問題 f(i-1) ：將 src 頂部 i-1 個圓盤藉助 tar 移到 buf
	dfsHanota(i-1, src, tar, buf)
	// 子問題 f(1) ：將 src 剩餘一個圓盤移到 tar
	move(src, tar)

	fmt.Println("圆盘移动后：")
	fmt.Print("A = ")
	PrintList(src)
	fmt.Print("B = ")
	PrintList(buf)
	fmt.Print("C = ")
	PrintList(tar)

	// 子問題 f(i-1) ：將 buf 頂部 i-1 個圓盤藉助 src 移到 tar
	dfsHanota(i-1, buf, src, tar)
}

/* 求解汉诺塔问题 */
func solveHanota(A, B, C *list.List) {
	n := A.Len()
	// 将 A 顶部 n 个圆盘借助 B 移到 C
	dfsHanota(n, A, B, C)
}
