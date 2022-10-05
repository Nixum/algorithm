package greedy

import (
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	var arr1 []int
	var arr2 []int
	arr1 = []int{1,2,3}
	arr2 = []int{1,1}
	fmt.Println(findContentChildren(arr1, arr2))

	arr1 = []int{1,2}
	arr2 = []int{1,2,3}
	fmt.Println(findContentChildren(arr1, arr2))
}

func TestWiggleMaxLength(t *testing.T) {
	var arr []int
	arr = []int{1,7,4,9,2,5}
	fmt.Println(wiggleMaxLength(arr))

	arr = []int{1,17,5,10,13,15,10,5,16,8}
	fmt.Println(wiggleMaxLength(arr))

	arr = []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(wiggleMaxLength(arr))

	arr = []int{1,1}
	fmt.Println(wiggleMaxLength(arr))
}

func TestMaxSubArray(t *testing.T) {
	var arr []int
	arr = []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(arr))

	arr = []int{1}
	fmt.Println(maxSubArray(arr))

	arr = []int{5,4,-1,7,8}
	fmt.Println(maxSubArray(arr))
}

func TestCanJump(t *testing.T) {
	var arr []int
	arr = []int{2,3,1,1,4}
	fmt.Println(canJump(arr))

	arr = []int{1}
	fmt.Println(canJump(arr))

	arr = []int{3,2,1,0,4}
	fmt.Println(canJump(arr))
}

func TestJump(t *testing.T) {
	var arr []int
	arr = []int{2,3,1,1,4}
	fmt.Println(jump(arr))
}

func TestLargestSumAfterKNegations(t *testing.T) {
	var arr []int
	arr = []int{4,3,2}
	fmt.Println(largestSumAfterKNegations(arr, 1))

	arr = []int{3,-1,0,2}
	fmt.Println(largestSumAfterKNegations(arr, 3))

	arr = []int{2,-3,-1,5,-4}
	fmt.Println(largestSumAfterKNegations(arr, 2))
}