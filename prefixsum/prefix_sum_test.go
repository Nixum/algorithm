package prefixsum

import (
	"fmt"
	"testing"
)

func TestRangeSumQuery(t *testing.T) {
	arr := NumArrayConstructor2([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(arr.res)
	fmt.Println(arr.SumRange(0, 2))
	fmt.Println(arr.SumRange(2, 5))
	fmt.Println(arr.SumRange(0, 5))

	arr = NumArrayConstructor2([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(arr.res)
	fmt.Println(arr.SumRange(2, 4))
	fmt.Println(arr.SumRange(2, 3))
}

func TestSubarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

func TestRangeSumQuery2d(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	s := NumMatrixConstructor(matrix)
	fmt.Println(s)
	fmt.Println(s.SumRegion(2, 1, 4, 3))
	fmt.Println(s.SumRegion(1, 1, 2, 2))
	fmt.Println(s.SumRegion(1, 2, 2, 4))
	fmt.Println(s.SumRegion(2, 1, 4, 3))
	fmt.Println(s.SumRegion(1, 2, 2, 4))
}

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf2([]int{1,2,3,4}))
}