package dp

import "fmt"

// 字符串中的最长回文子串
// dp[i][j]：字符串s在[i, j]范围内是否是回文串
// 当s[i] == s[j], 分情况
// i == j, dp[i][j] = true,
// i + 1 == j, dp[i][j] = true
// j - i > 1 && dp[i+1][j-1] == true, dp[i][j] = true
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j - i <= 1 {
					dp[i][j] = true
					if j - i + 1 > len(res) {
						res = s[i:j+1]
					}
				} else {
					dp[i][j] = dp[i + 1][j - 1]
					if dp[i][j] && j - i + 1 > len(res) {
						res = s[i:j+1]
					}
				}
			}
		}
	}
	fmt.Println(dp)
	return res
}
