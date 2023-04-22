package monotonic_stack

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	nums := []int{}
	nums = []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(nums))

	nums = []int{4,2,0,3,2,5}
	fmt.Println(trap(nums))
}

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(calcInMaximalRectangle([]int{2,1,5,6,2,3}))
}

func TestGetMidNum(t *testing.T) {
	fmt.Println(getMidNum2([]int{2,1,3,4,5,7,6}))
	fmt.Println(getMidNum2([]int{3,2,1,4,7,6,5}))
	fmt.Println(getMidNum2([]int{7,6,5,4,3,2,1}))
	fmt.Println(getMidNum2([]int{6,5,4,3,2,1,7}))
	fmt.Println(getMidNum2([]int{1,7,6,5,4,3,2}))
}