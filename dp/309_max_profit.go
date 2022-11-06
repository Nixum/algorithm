package dp

import "algorithm/common"

// 卖出后的一天是冷冻期
// dp[i][0]: 第i天买入股票所持最多现金，有场景1：前一天已经买入，场景2：前一天是冷冻期,今天买入
// dp[i][1]: 第i天卖出股票所持最多现金，有场景1：前一天已经卖出，场景2：前一天买入，今天卖出
// dp[i][2]: 第i天卖出股票冷冻期所持最多现金，有场景1：前一天卖出
func maxProfitV(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = common.Max(dp[i - 1][0], dp[i - 1][2] - prices[i])
		dp[i][1] = common.Max(dp[i - 1][1], dp[i - 1][0] + prices[i])
		dp[i][2] = dp[i - 1][1]
	}
	return common.Max(dp[len(prices) - 1][1], dp[len(prices) - 1][2])
}