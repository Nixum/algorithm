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

func TestMinEatingSpeed(t *testing.T) {
	fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6))
}

func TestShipWithinDays(t *testing.T) {
	fmt.Println(shipWithinDays([]int{1,2,3,4,5,6,7,8,9,10}, 5))
	fmt.Println(shipWithinDays([]int{3,2,2,4,1,4}, 3))
	fmt.Println(shipWithinDays([]int{1,2,3,1,1}, 4))
}

func TestPreimageSizeFZF(t *testing.T) {
	fmt.Println(preimageSizeFZF(0))
	fmt.Println(preimageSizeFZF(5))
	fmt.Println(preimageSizeFZF(3))
	fmt.Println(preimageSizeFZF(1000000000))
}

func TestSearchMatrix2(t *testing.T) {
	fmt.Println(searchMatrix2([][]int{{1,3}}, 3))
}