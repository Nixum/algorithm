package dp

import "algorithm/common"

// 先把题目转换
// 因为一定有组合 left - right = target，
// --》 left - (sum - right) = target, 因为sum和target是固定的，
// 所以 left = (target + sum) / 2，
// 题目就变成了求和为left有多少种方式了
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if common.Abs(target) > sum {
		return 0
	}
	if (sum + target) % 2 == 1 {
		return 0
	}
	bagSize := target + (sum - target) / 2
	dp := make([]int, bagSize + 1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			// 这种是装满容量j的背包，最多是装多少，
			// 但题目求的是有多少种装的方式
			//dp[j] = common.Max(dp[j], dp[j - nums[i]] + nums[i])
			dp[j] += dp[j - nums[i]]
		}
	}
	return dp[bagSize]
}