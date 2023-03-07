package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var arr []int
	arr = []int{1,3,5,2,6,8,3,5,7,0,9}
	quickSort(arr)
	fmt.Println(arr)

	arr = []int{1,3,5,2,6,8,3,5,7,0}
	quickSort(arr)
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
	var arr []int
	arr = []int{1,3,5,2,6,8,3,5,7,0,9}
	heapSort(arr)
	fmt.Println(arr)

	arr = []int{1,3,5,2,6,8,3,5,7,0}
	heapSort(arr)
	fmt.Println(arr)
}

func TestFindKthLargest(t *testing.T) {
	var arr []int
	arr = []int{3,1,2,4}
	fmt.Println(findKthLargest(arr, 2))

	arr = []int{3,2,3,1,2,4,5,5,6}
	fmt.Println(findKthLargest(arr, 4))
}

func TestMergeSort(t *testing.T) {
	var arr []int
	arr = []int{1,3,5,2,6,8,3,5,7,0,9}
	mergeSort(arr)
	fmt.Println(arr)

	arr = []int{1,3,5,2,6,8,3,5,7,0}
	mergeSort(arr)
	fmt.Println(arr)
}