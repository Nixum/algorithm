package dp

import "algorithm/common"

func maxProfitVI(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = common.Max(dp[i - 1][0], dp[i - 1][1] - prices[i])
		dp[i][1] = common.Max(dp[i - 1][1], dp[i - 1][0] + prices[i] - fee)
	}
	return dp[len(prices) - 1][1]
}

// 贪心解法
func maxProfitVI2(prices []int, fee int) int {
	    res := 0
	    minPrice := prices[0]
	    for i := 1; i < len(prices); i++ {
	        if minPrice > prices[i] {
	            minPrice = prices[i]
	        }
	        if prices[i] >= minPrice && prices[i] <= minPrice + fee {
	            continue
	        }
	        if prices[i] > minPrice + fee {
	            res += prices[i] - minPrice - fee
				// 因为这里判断的是只要大于最低金额和手续费就往结果里加，
				// 不一定是真正的卖出，
	            minPrice = prices[i] - fee
	        }
	    }
	    return res
}