package greedy

import "algorithm/common"

func candy(nums []int) int {
	candy := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i - 1] {
			candy[i] += candy[i - 1] + 1
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i + 1] {
			// 注意这里是等于，因为要求最少，而左边已经算过了，所以max的时候是candy[i]
			candy[i] = common.Max(candy[i], candy[i + 1] + 1)
		}
	}

	res := 0
	for i := 0; i < len(candy); i++ {
		res += candy[i] + 1
	}
	return res
}