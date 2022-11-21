package dp

import "algorithm/common"

// dp[i][j]: 前i-1个字符串转换成前j-1个字符串所需操作的最少次数
func editDistance(w1 string, w2 string) int {
	dp := make([][]int, len(w1) + 1)
	for i := 0; i <= len(w1); i++ {
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
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				// word1删除元素
				// word1添加元素，相当于word2删除元素
				// word1替换元素，替换word1[i-1], 让word1[i-1]==word2[j-1]，相当于以word1[i-2]和word2[j-2]结尾的字符串的最小编辑次数+1个替换操作
				dp[i][j] = common.Min(common.Min(dp[i - 1][j] + 1, dp[i][j - 1] + 1), dp[i - 1][j - 1] + 1)
			}
		}
	}
	return dp[len(w1)][len(w2)]
}
