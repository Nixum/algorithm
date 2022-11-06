package dp

import (
	"algorithm/common"
	"math"
)

// 因为只能买卖一次
func maxProfit(prices []int) int {
	low := math.MaxInt64
	res := 0
	for i := 0; i < len(prices); i++ {
		low = common.Min(low, prices[i])
		res = common.Max(res, prices[i] - low)
	}
	return res
}

// dp解法
// dp[i][0]: 第i天持有股票所得最多现金
// dp[i][1]：第i天不持有股票所得最多现金
func maxProfit1(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		// 因为只能买一次，所以是 -prices[i]
		dp[i][0] = common.Max(dp[i - 1][0], -prices[i])
		dp[i][1] = common.Max(dp[i - 1][1], dp[i - 1][0] + prices[i])
	}
	return dp[len(prices) - 1][1]
}