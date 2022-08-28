package stack
type MyQueue struct {
	stack []int
	back []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

func (q *MyQueue) Push(x int)  {
	for i := len(q.back) - 1; i >= 0; i-- {
		q.stack = append(q.stack, q.back[i])
	}
	q.back = q.back[:0]
	q.stack = append(q.stack, x)
}

func (q *MyQueue) Pop() int {
	for i := len(q.stack) - 1; i >= 0; i-- {
		q.back = append(q.back, q.stack[i])
	}
	q.stack = q.stack[:0]
	if len(q.back) == 0 {
		return -1
	}
	res := q.back[len(q.back) - 1]
	q.back = q.back[:len(q.back) - 1]
	return res
}

func (q *MyQueue) Peek() int {
	res := q.Pop()
	if res == -1 {
		return res
	}
	q.back = append(q.back, res)
	return res
}

func (q *MyQueue) Empty() bool {
	if len(q.stack) == 0 && len(q.back) == 0 {
		return true
	}
	return false
}


// --------------------------------
type MyQueue1 struct {
	stack1 []int // 只进
	stack2 []int // 只出
}

func MyQueueConstructor1() MyQueue1 {
	return MyQueue1{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (q *MyQueue1) Push(x int)  {
	q.stack1 = append(q.stack1, x)
}

func (q *MyQueue1) Pop() int {
	if len(q.stack2) > 0 {
		res := q.stack2[len(q.stack2) - 1]
		q.stack2 = q.stack2[:len(q.stack2) - 1]
		return res
	}
	if len(q.stack1) > 0 && len(q.stack2) == 0 {
		for i := len(q.stack1) - 1; i >= 0; i-- {
			if i == 0 {
				res := q.stack1[i]
				q.stack1 = q.stack1[:0]
				return res
			}
			q.stack2 = append(q.stack2, q.stack1[i])
		}
	}
	return -1
}

func (q *MyQueue1) Peek() int {
	if len(q.stack2) > 0 {
		res := q.stack2[len(q.stack2) - 1]
		return res
	}
	if len(q.stack1) > 0 && len(q.stack2) == 0 {
		for i := len(q.stack1) - 1; i >= 0; i-- {
			if i == 0 {
				res := q.stack1[i]
				q.stack2 = append(q.stack2, q.stack1[i])
				q.stack1 = q.stack1[:0]
				return res
			}
			q.stack2 = append(q.stack2, q.stack1[i])
		}
	}
	return -1
}

func (q *MyQueue1) Empty() bool {
	if len(q.stack1) == 0 && len(q.stack2) == 0 {
		return true
	}
	return false
}
