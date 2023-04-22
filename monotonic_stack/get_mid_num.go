package monotonic_stack

import (
	"algorithm/common"
	"math"
)

// 给出一个元素无序的数组，求出一个数，使得其左边的数都小于它，右边的数都大于等于它。
// 思路：使用单调栈，栈内递增，如果遇到比较小的，pop掉栈顶元素，实现右边的值都比该数大
// 再维护一个max值，比max值大才能进栈，实现左边的值都比该数小
func getMidNum(nums []int) []int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if nums[i] > max {
			stack = append(stack, i)
		}
		max = common.Max(max, nums[i])
	}
	return stack
}

// 另一种思路
// leftMax 记录nums[i]左边的最大值, rightMin 记录nums[i]右边的最小值
// 当 leftMax[i]<nums[i]<=rightMin[i]，就符合题意
func getMidNum2(nums []int) []int {
	res := make([]int, 0)
	leftMax := make([]int, len(nums))
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		leftMax[i] = max
	}
	rightMin := make([]int, len(nums))
	min := math.MaxInt64
	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] < min {
			min = nums[i]
		}
		rightMin[i] = min
	}
	for i := 0; i < len(nums); i++ {
		if leftMax[i] <= nums[i] && nums[i] <= rightMin[i] {
			res = append(res, i)
		}
	}
	return res
}