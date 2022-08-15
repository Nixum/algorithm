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