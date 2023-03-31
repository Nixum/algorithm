package dp

// dp[i] 表示以s[i]结尾的最长有效括号
// 如果s[i] == '(', 此时 dp[i] = 0
// 如果s[i] == ')', 要分情况分析
// 		如果s[i-1] == '(', 此时dp[i] = dp[i-2] + 2
// 		如果s[i-1] == ')', 此时又分两种情况
// 			如果s[i-dp[i-1]-1] == '(', 此时dp[i] = dp[i-1]+2
// 			因为s[i-dp[i-1]-2]之前有可能已经是有效括号了，所以还要加上 dp[i-dp[i-1]-2]
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	max := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i - 2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 {
						dp[i] += dp[i-dp[i-1]-2]
					}
				}
			}
			if max < dp[i] {
				max = dp[i]
			}
		}
	}
	return max
}
