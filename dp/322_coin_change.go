package dp

import (
	"algorithm/common"
	"math"
)

// dp[j]: 凑齐面值为j时所需要的最少硬币数量为dp[j]
// dp[j] = min(dp[j], dp[j - coins[i]] + 1)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	dp[0] = 0
	// 求最小，所以默认都是最大值
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = common.Min(dp[j], dp[j - coins[i]] + 1)
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}
