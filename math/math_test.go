package math

import (
	"fmt"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	fmt.Println(trailingZeroes(125))
}

func TestMedianFinder(t *testing.T) {
	//desc := &priorityQueue{
	//	isDesc: true,
	//	nums:   make([]int, 0),
	//}
	//heap.Push(desc, 3)
	//heap.Push(desc, 4)
	//heap.Push(desc, 1)
	//heap.Push(desc, 2)
	//fmt.Println(heap.Pop(desc))
	//fmt.Println(heap.Pop(desc))
	//fmt.Println(heap.Pop(desc))
	//fmt.Println(heap.Pop(desc))

	f := MedianFinderConstructor()
	f.AddNum(2)
	f.AddNum(1)
	fmt.Println(f.FindMedian())
	f.AddNum(3)
	fmt.Println(f.FindMedian())
}
