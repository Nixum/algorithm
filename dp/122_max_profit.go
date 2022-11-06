package dp

import "algorithm/common"

// 允许多次买卖
// dp[i][0] 第i天持有股票时最多现金
// dp[i][1] 第i天不持有股票时的最多现金
func maxProfitII(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = common.Min(dp[i - 1][0], dp[i - 1][1] - prices[i])
		dp[i][1] = common.Max(dp[i - 1][1], dp[i - 1][0] + prices[i])
	}
	return dp[len(prices) - 1][1]
}
