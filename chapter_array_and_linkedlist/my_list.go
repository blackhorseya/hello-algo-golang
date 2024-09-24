package chapter_array_and_linkedlist

type myList struct {
	arrCapacity int
	arr         []int
	arrSize     int
	extendRatio int
}

func newMyList() *myList {
	return &myList{
		arrCapacity: 10,
		arr:         make([]int, 10),
		arrSize:     0,
		extendRatio: 2,
	}
}

func (l *myList) size() int {
	return l.arrSize
}

func (l *myList) capacity() int {
	return l.arrCapacity
}

// TODO: 2024/9/24|sean|implement the following methods
