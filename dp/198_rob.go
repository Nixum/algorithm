package dp

import "algorithm/common"

// dp[i]: 打劫前i家能偷窃的最大金额
// 有两种场景：
// 偷i房：dp[i] = dp[i - 2] + nums[i]
// 不偷i房：dp[i] = dp[i - 1]
// 所以，dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
func robI(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = common.Max(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = common.Max(dp[i - 1], dp[i - 2] + nums[i])
	}
	return dp[len(nums) - 1]
}
