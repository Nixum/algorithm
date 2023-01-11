package monotonic_stack

import "algorithm/common"

// 利用单调递增栈，入栈时，如果栈顶元素大于当前值，则出栈，直到栈顶元素小于当前值，当前值入栈
// 利用单调栈的特性，通过单调栈找到第一个比当前值小的值，说明这个值是右边值的第一个小值，就可以计算右边这个值的面积大小了
// 所以触发时机就是当前值小于栈顶元素，此时的栈顶元素左边的值就是该栈顶元素的第一个小值，此时就能计算栈顶元素的面积，直到栈顶元素小于当前值
func largestRectangleArea(heights []int) int {
	// 单调栈 - 递增，存的是数组的下标
	// 关键，需要在末尾加0，确保在完全递增的时候会进行计算
	heights = append(heights, 0)
	stack := make([]int, 0)
	preIndex := 0
	res := heights[0]
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[top(stack)] {
			stack, preIndex = pop(stack)
			prePreIndex := top(stack) + 1
			res = common.Max(res, (i - prePreIndex) * heights[preIndex])
		}
		stack = append(stack, i)
	}
	return res
}

// 思路是对的，但是超时了，时间复杂度是O(n^2)
// 原理是找出以i为最矮的柱子，找出左右两边第一个比heights[i]小的值，下标相减作为长，最小值作为宽
func largestRectangleArea2(heights []int) int {
	res := -1
	for i := 0; i < len(heights); i++ {
		left := i
		right := i
		for left >= 0 && heights[left] >= heights[i] {
			left--
		}
		for right < len(heights) && heights[right] >= heights[i] {
			right++
		}
		res = common.Max(res, common.Min(heights[right], heights[left]) * (right - left - 1))
	}
	return res
}

// 错误的思路
func largestRectangleArea3(heights []int) int {
	descStack := make([]int, 0)
	res := heights[0]
	var t int
	descStack = append(descStack, heights[0])
	width := 1
	for i := 1; i < len(heights); i++ {
		descStack, t = pop(descStack)
		if heights[i] <= t {
			width++
			descStack = append(descStack, t)
			descStack = append(descStack, heights[i])
			res = common.Max(res, heights[i] * width)
		} else if heights[i] == 0 {
			width = 0
			descStack = make([]int, 0)
		} else {
			width++
			if heights[i] > res {
				width = 1
				descStack = []int{heights[i]}
				res = heights[i]
			} else {
				descStack = append(descStack, t)
				res = common.Max(res, t * width)
			}
		}
	}
	return res
}
