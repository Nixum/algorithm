package dp

import "algorithm/common"

// dp[i][j], 前i-1个字符的与前j-1个字符，要达到相同时所需删除元素的最小个数
func minDistance(w1 string, w2 string) int {
	dp := make([][]int, len(w1) + 1)
	for i := 0; i <= len(w2); i++ {
		dp[i] = make([]int, len(w2) + 1)
	}
	for i := 0; i <= len(w1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(w2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(w1); i++ {
		for j := 1; j <= len(w2); j++ {
			if w1[i - 1] == w2[j - 1] {
				// 不用删，所以拿之前的结果即可
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				// 三种情况，取最小
				// 1：删除w1[i-1]，即dp[i-1][j] + 1
				// 2：删除w2[j-1]，即dp[i][j-1] + 1
				// 3: 即删除w1[i-1]和w2[j-1], 即dp[i-1][j-1] + 2,
				// 因为dp[i-1][j-1] + 1等于dp[i-1][j]或dp[i][j-1]，
				// 所以递推公式可简化为：dp[i][j] = min(dp[i - 1][j] + 1, dp[i][j - 1] + 1);
				dp[i][j] = common.Min(dp[i - 1][j] + 1, dp[i][j - 1] + 1)
			}
		}
	}
	return dp[len(w1)][len(w2)]
}
