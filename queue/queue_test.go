package queue

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1,1,1,2,2,3,3,3,3}
	fmt.Println(topKFrequent(nums, 2))
}

func TestMyStack(t *testing.T) {
	s := MyStackConstructor()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}
