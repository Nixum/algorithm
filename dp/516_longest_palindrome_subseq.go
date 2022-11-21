package dp

import "algorithm/common"

// 字符串中的最长回文子序列
// dp[i][j]：字符串s在[i, j]范围内最长的回文子序列的长度
// 当s[i] == s[j], dp[i][j] = dp[i+1][j-1] + 2
// 当s[i] != s[j], dp[i][j] = dp[i+1][j-1]
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = common.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s) - 1]
}