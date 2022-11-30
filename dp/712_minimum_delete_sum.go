package dp

import (
	"algorithm/common"
	"math"
)

// dp[i][j]: s1前i-1个字符和s2前j-1个字符使其相等时的最少删除的字符的ascII和
func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1) + 1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2) + 1)
		for j := 0; j <= len(s2); j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i - 1][0] + int(s1[i-1])
	}
	for i := 1; i <= len(s2); i++ {
		dp[0][i] = dp[0][i - 1] + int(s2[i-1])
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i - 1] == s2[j - 1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.Min(common.Min(dp[i-1][j] + int(s1[i-1]), dp[i][j-1] + int(s2[j-1])), dp[i-1][j-1] + int(s1[i-1]) + int(s2[j-1]))
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
