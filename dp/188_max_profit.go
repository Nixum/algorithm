package dp

import "algorithm/common"

func maxProfitIV(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, k*2 + 1)
	}
	for i := 1; i < 2 * k + 1; i += 2 {
		dp[0][i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j < 2 * k + 1; j++ {
			p := prices[i]
			if j % 2 == 1 {
				p = -p
			}
			dp[i][j] = common.Max(dp[i - 1][j], dp[i - 1][j - 1] + p)
		}
	}
	return dp[len(prices) - 1][2 * k]
}
