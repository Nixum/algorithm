package dp

// dp[i][j] 前i-1个字符的s是否匹配前j-1个字符的p
// 分情况讨论：
// s[i-1] == p[j-1] || p[j-1] == '.', dp[i][j] = dp[i-1][j-1]
// s[i-1] != p[j-1], p[j-1] != '*', dp[i][j] = false
// s[i-1] != p[j-1], p[j-1] == '*', 再分情况
// 	  s[i-1] == p[j-2] || p[j-2] == '.', dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j], 分别对应匹配0次，匹配一次，匹配多次
//    s[i-1] != p[j-2], dp[i][j] = dp[i][j-2]
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if p[j-1] == '*' {
					if s[i-1] == p[j-2] || p[j-2] == '.' {
						dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
					} else {
						dp[i][j] = dp[i][j-2]
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
