package chapter_stack_and_queue

type arrayStack struct {
	data []int
}

func newArrayStack() *arrayStack {
	return &arrayStack{
		// 初始化一个容量为16的数组
		data: make([]int, 0, 16),
	}
}

func (s *arrayStack) size() int {
	return len(s.data)
}

func (s *arrayStack) isEmpty() bool {
	return s.size() == 0
}

func (s *arrayStack) push(v int) {
	s.data = append(s.data, v)
}

func (s *arrayStack) peak() any {
	if s.isEmpty() {
		return nil
	}

	return s.data[s.size()-1]
}

func (s *arrayStack) pop() any {
	if s.isEmpty() {
		return nil
	}

	v := s.data[s.size()-1]
	s.data = s.data[:s.size()-1]
	return v
}

/* 獲取 Slice 用於列印 */
func (s *arrayStack) toSlice() []int {
	return s.data
}
