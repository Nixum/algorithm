package math

import "container/heap"

type MedianFinder struct {
	descHeap *priorityQueue
	ascHeap  *priorityQueue
}

func MedianFinderConstructor() MedianFinder {
	asc := &priorityQueue{
		isDesc: false,
		nums:   make([]int, 0),
	}
	heap.Init(asc)
	desc := &priorityQueue{
		isDesc: true,
		nums:   make([]int, 0),
	}
	heap.Init(desc)
	return MedianFinder{
		ascHeap:  asc,
		descHeap: desc,
	}
}

// 还要维护降序堆的堆顶元素大于升序堆
// 所以add的时候要先去一边的堆里走一遭，进行置换，再加入到另一个堆
func (m *MedianFinder) AddNum(num int)  {
	// 哪边少加哪边
	if m.ascHeap.Len() >= m.descHeap.Len() {
		heap.Push(m.ascHeap, num)
		heap.Push(m.descHeap, heap.Pop(m.ascHeap))
	} else {
		heap.Push(m.descHeap, num)
		heap.Push(m.ascHeap, heap.Pop(m.descHeap))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.ascHeap.Len() > m.descHeap.Len() {
		return float64(m.ascHeap.Peek())
	} else if m.ascHeap.Len() < m.descHeap.Len() {
		return float64(m.descHeap.Peek())
	}
	return float64(m.ascHeap.Peek() + m.descHeap.Peek()) / 2.0
}

// 优先级队列 --------------------------
type priorityQueue struct {
	isDesc bool
	nums []int
}

func (q *priorityQueue) Len() int {
	return len(q.nums)
}

func (q *priorityQueue) Less(i, j int) bool {
	if q.isDesc {
		return q.nums[i] > q.nums[j]
	}
	return q.nums[i] < q.nums[j]
}

func (q *priorityQueue) Swap(i, j int) {
	q.nums[i], q.nums[j] = q.nums[j], q.nums[i]
}

func (q *priorityQueue) Push(x interface{}) {
	q.nums = append(q.nums, x.(int))
}

func (q *priorityQueue) Pop() interface{} {
	old := q.nums
	n := len(old)
	x := old[n-1]
	q.nums = old[0 : n-1]
	return x
}

func (q *priorityQueue) Peek() int {
	return q.nums[0]
}
