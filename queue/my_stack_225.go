package queue

// ----------------------
type MyStack struct {
	// 只用一个队列来周转
	queue []int
}


func MyStackConstructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (s *MyStack) Push(x int)  {
	s.queue = append(s.queue, x)
}

func (s *MyStack) Pop() int {
	if len(s.queue) == 0 {
		return -1
	}
	if len(s.queue) == 1 {
		res := s.queue[0]
		s.queue = s.queue[0:0]
		return res
	}
	n := len(s.queue) - 1
	res := s.queue[n]
	for i := 0; i < n; i++ {
		s.queue = append(s.queue, s.queue[i])
	}
	s.queue = s.queue[n+1:]
	return res
}

func (s *MyStack) Top() int {
	res := s.Pop()
	s.Push(res)
	return res
}


func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}


// -----------------------
type MyStack1 struct {
	// 两个slice得模拟队列，所以不能用一个slice实现
	left []int
	right []int
}

func MyStackConstructor1() MyStack1 {
	return MyStack1{
		left: make([]int, 0),
		right: make([]int, 0),
	}
}

func (s *MyStack1) Push(x int)  {
	if len(s.left) != 0 {
		s.left = append(s.left, x)
		return
	}
	if len(s.right) != 0 {
		s.right = append(s.right, x)
		return
	}
	s.left = append(s.left, x)
}

func (s *MyStack1) Pop() int {
	if len(s.left) > 0 {
		tmp := s.right
		s.right = s.left[0:len(s.left) - 1]
		res := s.left[len(s.left) - 1]
		s.left = tmp
		return res
	}
	if len(s.right) > 0 {
		tmp := s.left
		s.left = s.right[0:len(s.right) - 1]
		res := s.right[len(s.right) - 1]
		s.right = tmp
		return res
	}
	return -1
}

func (s *MyStack1) Top() int {
	if len(s.left) == 0 && len(s.right) == 0 {
		return -1
	}
	res := s.Pop()
	s.Push(res)
	return res
}


func (s *MyStack1) Empty() bool {
	if len(s.right) == 0 && len(s.left) == 0 {
		return true
	}
	return false
}
