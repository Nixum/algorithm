package greedy

import "algorithm/common"

func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = common.Max(cover, i + nums[cover])
		if cover >= len(nums) - 1 {
			return true
		}
	}
	return false
}
