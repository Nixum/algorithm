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

func TestHammingWeight(t *testing.T) {
	fmt.Println(hammingWeight(11))
}

func TestIsPowerOfTwo(t *testing.T) {
	fmt.Println(isPowerOfTwo(11))
	fmt.Println(isPowerOfTwo(8))
}

func TestPreimageSizeFZF(t *testing.T) {
	fmt.Println(preimageSizeFZF(3))
}

func TestMyPow(t *testing.T) {
	fmt.Println(myPow2(2, 10))
}

func TestMySqrt(t *testing.T) {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}

func TestGenerate(t *testing.T) {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}

func TestFindMedianSortedArrays(t *testing.T) {
	arr1 := []int{}
	arr2 := []int{3}
	fmt.Println(findMedianSortedArrays(arr1, arr2))
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(19))
}
