package double_point

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3,2,2,3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)

	nums = []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)

	nums = []int{4,2,0,2,2,1,4,4,1,4,3,2}
	fmt.Println(removeElement(nums, 4))
	fmt.Println(nums)
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{0,1,2,2,3,3,4,4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func TestMoveZero(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1,0}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0, 1, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func TestBackspaceCompare(t *testing.T) {
	fmt.Println(backspaceCompare("ab#c", "ac#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#c", "b"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
}

func TestSortedSquares(t *testing.T) {
	fmt.Println(sortedSquares2([]int{-4,-1,0,3,10}))
}

func TestTrap(t *testing.T) {
	fmt.Println(trap2([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{1,1,-2}))
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum([]int{0,1,1}))
	fmt.Println(threeSum([]int{0,0,0}))
	fmt.Println(threeSum([]int{0,0,0,0,0}))
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2,7,11,15,9,0,-2}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}
