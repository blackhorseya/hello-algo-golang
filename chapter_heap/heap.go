package chapter_heap

// Go 語言中可以透過實現 heap.Interface 來構建整數大頂堆積
// 實現 heap.Interface 需要同時實現 sort.Interface
type intHeap []any

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	// 如果實現小頂堆積，則需要調整為小於號
	return (*h)[i].(int) > (*h)[j].(int)
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(x any) {
	// Push 和 Pop 使用 pointer receiver 作為參數
	// 因為它們不僅會對切片的內容進行調整，還會修改切片的長度。
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Top 获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}
