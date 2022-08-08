package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	fmt.Println(BinarySearch([]int{-1,0,3,5,9,12}, 9))
	fmt.Println(BinarySearch([]int{-1,0,3,5,9,12}, 2))

	fmt.Println(BinarySearch2([]int{-1,0,3,5,9,12}, 9))
	fmt.Println(BinarySearch2([]int{-1,0,3,5,9,12}, 2))
}

func TestBinaryInsert(t *testing.T) {
	fmt.Println(SearchInsert([]int{1,3,5,6}, 2))
}

func TestMySqrt2(t *testing.T) {
	fmt.Println(MySqrt2(6))
}
