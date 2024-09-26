package chapter_heap

import (
	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

type maxHeap struct {
	// 使用切片而非数组，这样无须考虑扩容问题
	data []any
}

func newHeap() *maxHeap {
	return &maxHeap{
		data: make([]any, 0),
	}
}

/* 构造函数，根据切片建堆 */
func newMaxHeap(nums []any) *maxHeap {
	// 将列表元素原封不动添加进堆
	h := &maxHeap{data: nums}
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		// 堆化除叶节点以外的其他所有节点
		h.siftDown(i)
	}
	return h
}

func (h *maxHeap) left(i int) int {
	return 2*i + 1
}

func (h *maxHeap) right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

/* 获取堆大小 */
func (h *maxHeap) size() int {
	return len(h.data)
}

/* 判断堆是否为空 */
func (h *maxHeap) isEmpty() bool {
	return len(h.data) == 0
}

/* 訪問堆積頂元素 */
func (h *maxHeap) peek() any {
	if h.isEmpty() {
		return nil
	}
	return h.data[0]
}

/* 元素入堆積 */
func (h *maxHeap) push(val any) {
	// 先将元素添加到最后
	h.data = append(h.data, val)
	// 然后堆化
	h.siftUp(h.size() - 1)
}

/* 从节点 i 开始，从底至顶堆化 */
func (h *maxHeap) siftUp(i int) {
	for {
		parent := h.parent(i)
		if parent < 0 || h.data[i].(int) <= h.data[parent].(int) {
			break
		}
		h.swap(i, parent)
		i = parent
	}
}

/* 交换元素 */
func (h *maxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

/* 元素出堆 */
func (h *maxHeap) pop() any {
	if h.isEmpty() {
		return nil
	}
	// 交换堆顶元素和最后一个元素
	h.swap(0, h.size()-1)
	// 删除最后一个元素
	val := h.data[h.size()-1]
	h.data = h.data[:h.size()-1]
	// 堆化
	h.siftDown(0)
	return val
}

func (h *maxHeap) siftDown(i int) {
	for {
		// 获取左右节点 (left, right) 和最大节点 (largest)
		left, right, largest := h.left(i), h.right(i), i

		if left < h.size() && h.data[left].(int) > h.data[largest].(int) {
			largest = left
		}
		if right < h.size() && h.data[right].(int) > h.data[largest].(int) {
			largest = right
		}
		if largest == i {
			break
		}
		h.swap(i, largest)
		i = largest
	}
}

/* 打印堆（二叉树） */
func (h *maxHeap) print() {
	PrintHeap(h.data)
}
