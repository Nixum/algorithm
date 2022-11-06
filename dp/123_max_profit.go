package dp

import "algorithm/common"

// 最多买卖两次
// dp[i][0]：第i天不操作
// dp[i][1]：第i天第一次买入后所持有的现金
// dp[i][2]：第i天第一次卖出后所持有的现金
// dp[i][3]：第i天第二次买入后所持有的现金
// dp[i][4]：第i天第二次卖出后所持有的现金
func maxProfitIII(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		// max（前i-1天没有操作，前i-1天买入/卖出股票）
		dp[i][1] = common.Max(dp[i - 1][1], dp[i - 1][0] - prices[i])
		dp[i][2] = common.Max(dp[i - 1][2], dp[i - 1][1] + prices[i])
		dp[i][3] = common.Max(dp[i - 1][3], dp[i - 1][2] - prices[i])
		dp[i][4] = common.Max(dp[i - 1][4], dp[i - 1][3] + prices[i])
	}
	return dp[len(prices) - 1][4]
}