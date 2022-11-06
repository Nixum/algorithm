package dp

import "algorithm/common"

// dp[i]：打劫前i家能偷得最大金额
// 分成三种情况：
// 1. 不考虑首尾元素
// 2. 不考虑首元素，只考虑尾元素
// 3. 只考虑首元素，不考虑尾元素
// 实际上2、3已经包含了1
func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res := rangeInRobII(nums, 1, len(nums) - 2)
	res1 := rangeInRobII(nums, 0, len(nums) - 2)
	res2 := rangeInRobII(nums, 1, len(nums) - 1)
	return common.Max(common.Max(res, res1), res2)
}

func rangeInRobII(nums []int, start, end int) int {
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	if start + 1 <= end {
		dp[start + 1] = common.Max(nums[start + 1], nums[start])
	}
	for i := start + 2; i <= end; i++ {
		dp[i] = common.Max(dp[i - 1], dp[i - 2] + nums[i])
	}
	return dp[end]
}
