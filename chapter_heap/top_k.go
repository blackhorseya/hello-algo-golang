package chapter_heap

import (
	"container/heap"
)

type minHeap []any

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i].(int) < (*h)[j].(int) }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

// Push heap.Interface 的方法，实现推入元素到堆
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop heap.Interface 的方法，实现弹出堆顶元素
func (h *minHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Top 获取堆顶元素
func (h *minHeap) Top() any {
	return (*h)[0]
}

/* 基于堆查找数组中最大的 k 个元素 */
func topKHeap(nums []int, k int) *minHeap {
	// 初始化一个小顶堆
	h := &minHeap{}
	heap.Init(h)

	// 先将前 k 个元素推入堆中
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	// 遍历剩余元素, 若大于堆顶元素则替换
	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0].(int) {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	return h
}
