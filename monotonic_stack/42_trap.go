package monotonic_stack

import (
	"algorithm/common"
)

func trap(height []int) int {
	res := 0
	stack := []int{0}
	var cur = 0
	for i := 1; i < len(height); i++ {
		if len(stack) > 0 && height[i] == height[top(stack)] {
			stack, _ = pop(stack)
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && height[i] > height[top(stack)] {
			stack, cur = pop(stack)
			pre := top(stack)
			if pre == -1 {
				continue
			}
			res += (i - pre - 1) * (common.Min(height[pre], height[i]) - height[cur])
		}
		stack = append(stack, i)
	}
	return res
}
