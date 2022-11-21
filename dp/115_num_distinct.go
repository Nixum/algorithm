package dp

// dp[i][j]：前i-1的字符串s，在前j-1的字符串t中出现个个数
func numDistinct(s string, t string) int {
	dp := make([][]int, len(s) + 1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t) + 1)
	}
	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}
	for j := 0; j <= len(t); j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i - 1] == t[j - 1] {
				// 还要 加上 dp[i - 1][j] 是要匹配字符有重复的场景
				dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j]
			} else {
				// 不用s[i - 1]来匹配
				dp[i][j] = dp[i - 1][j]
			}
		}
	}
	return dp[len(t)][len(s)]
}