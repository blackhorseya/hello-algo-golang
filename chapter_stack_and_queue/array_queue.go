package chapter_stack_and_queue

type arrayQueue struct {
	nums          []int // 用于存储元素的数组
	front         int   // 队首指针
	queueSize     int   // 队列大小
	queueCapacity int   // 队列容量
}

func newArrayQueue(capacity int) *arrayQueue {
	return &arrayQueue{
		nums:          make([]int, capacity),
		front:         0,
		queueSize:     0,
		queueCapacity: capacity,
	}
}

func (q *arrayQueue) size() int {
	return q.queueSize
}

func (q *arrayQueue) isEmpty() bool {
	return q.size() == 0
}

func (q *arrayQueue) push(num int) {
	// 當 rear == queCapacity 表示佇列已滿
	if q.queueSize == q.queueCapacity {
		return
	}
	// 計算佇列尾指標，指向佇列尾索引 + 1
	// 透過取餘操作實現 rear 越過陣列尾部後回到頭部
	rear := (q.front + q.queueSize) % q.queueCapacity
	// 將 num 新增至佇列尾
	q.nums[rear] = num
	q.queueSize++
}

func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}

	return q.nums[q.front]
}

func (q *arrayQueue) pop() any {
	if q.isEmpty() {
		return nil
	}

	num := q.nums[q.front]
	// 將 front 指標指向下一個元素
	q.front = (q.front + 1) % q.queueCapacity
	q.queueSize--

	return num
}

/* 獲取 Slice 用於列印 */
func (q *arrayQueue) toSlice() []int {
	rear := q.front + q.queueSize
	if rear >= q.queueCapacity {
		rear %= q.queueCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}

	return q.nums[q.front:rear]
}
