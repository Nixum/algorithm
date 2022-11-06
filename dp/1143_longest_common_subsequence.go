package dp

import "algorithm/common"

// 因为边界问题，所以由i-1,j-1比较
// dp[i][j]: 前i-1个字符串与前j-1个字符串的最长公共子序列为dp[i][j]
// 与718的区别是，这里不在要求是连续的
// 递推公式：
// 如果t1[i - 1] == t2[j - 1], dp[i][j] = dp[i - 1][j - 1] + 1
// 否则, 从t1的结果推出，或者从t2的结果推出,
// dp[i][j] = Max(dp[i - 1][j], dp[i][j - 1])
func longestCommonSubsequence(t1 string, t2 string) int {
	dp := make([][]int, len(t1) + 1)
	for i := 0; i < len(t1) + 1; i++ {
		dp[i] = make([]int, len(t2) + 1)
	}
	for i := 1; i <= len(t1) ; i++ {
		for j := 1; j <= len(t2); j++ {
			if t1[i - 1] == t2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = common.Max(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}
	// 因为不要求是连续的，所以可以拿最后的结果
	return dp[len(t1)][len(t2)]
}
