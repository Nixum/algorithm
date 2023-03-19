package stack

import (
	"container/heap"
)

type MinStack struct {
	stack []int
	minStack []int
}


func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}


func (s *MinStack) Push(val int)  {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 || s.minStack[len(s.minStack)-1] >= val {
		s.minStack = append(s.minStack, val)
	}
}


func (s *MinStack) Pop()  {
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if s.minStack[len(s.minStack)-1] == val {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
}


func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}


func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

//
type MinStack1 struct {
	stack []int
	priority *priorityQ
}


func Constructor2() MinStack1 {
	return MinStack1{
		stack: make([]int, 0),
		priority:   &priorityQ{},
	}
}


func (s *MinStack1) Push(val int)  {
	s.stack = append(s.stack, val)
	heap.Push(s.priority, val)
}


func (s *MinStack1) Pop()  {
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.priority.PopVal(val)
}


func (s *MinStack1) Top() int {
	return s.stack[len(s.stack)-1]
}


func (s *MinStack1) GetMin() int {
	return s.priority.Top().(int)
}

//
type priorityQ []int

func (h priorityQ) Len() int           { return len(h) }
func (h priorityQ) Less(i, j int) bool { return h[i] < h[j] }
func (h priorityQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *priorityQ) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *priorityQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *priorityQ) Top() interface{} {
	return (*h)[0]
}

func (h *priorityQ) PopVal(val int) {
	for i := 0; i < h.Len(); i++ {
		if (*h)[i] == val {
			heap.Remove(h, i)
			break
		}
	}
	return
}
