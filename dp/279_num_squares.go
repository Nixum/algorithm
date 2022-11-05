package dp

import (
	"algorithm/common"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}
	for i := 1; i <= n / 2; i++ {
		for j := 1; j <= n; j++ {
			if j > i * i {
				dp[j] += common.Min(dp[j], dp[j - i * i] + 1)
			}
		}
	}
	return dp[n]
}
