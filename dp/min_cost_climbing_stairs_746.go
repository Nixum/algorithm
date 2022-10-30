package dp

import "algorithm/common"

// dp数组含义，到达第i个台阶所花费的最少体力为dp[i]
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = common.Min(dp[i - 1], dp[i - 2]) + cost[i]
	}
	// 最后一步不需要消耗，so在倒数第一和第二里二选一
	return common.Min(dp[len(dp) - 1], dp[len(dp) - 2])
}
