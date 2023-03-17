package double_point

import "algorithm/common"

func maxArea(height []int) int {
	left := 0
	right := len(height)-1
	leftMax := 0
	rightMax := len(height)-1
	res := 0
	for left < right {
		if height[leftMax] < height[left] {
			leftMax = left
		}
		if height[rightMax] < height[right] {
			rightMax = right
		}
		res = common.Max(res, (rightMax - leftMax) * common.Min(height[leftMax], height[rightMax]))
		if height[leftMax] < height[rightMax] {
			left++
		} else {
			right--
		}
	}
	return res
}
