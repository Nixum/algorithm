package slide_window

import (
	"math"
)
func minSubArrayLen(s int, nums []int) int {
	min := math.MaxInt64
	left := 0
	right := 0
	tmp := 0
	for right < len(nums) {
		tmp += nums[right]
		for tmp >= s {
			min = minN(min, right - left + 1)
			tmp -= nums[left]
			left++
		}
		right++
	}
	if min == math.MaxInt64 {
		return 0
	}
	return min
}

func minN(i, j int) int {
	if i > j {
		return j
	}
	return i
}