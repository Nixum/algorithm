package dp

import "fmt"

// s是否是t的子序列
// dp[i][j] 表示以下标i-1为结尾的字符串s，
// 和以下标j-1为结尾的字符串t，相同子序列的长度为dp[i][j]
// 所以最后一个字符一定是相同的
func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s) + 1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t) + 1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i - 1] == t[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				// 如果不等，说明t要把当前元素删掉
				// 此时看的是 s[i - 1] 与 t[j - 2]的结果
				dp[i][j] = dp[i][j - 1]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)][len(t)] == len(s)
}
