package dp

import "fmt"

// dp[i]，字符串的前i个字符可以拆分成一个或多个在字典种出现的单词
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool, 0)
	for i := range wordDict {
		dict[wordDict[i]] = true
	}
	dp := make([]bool, len(s) + 1)
	dp[0] = true
	// 先遍历背包
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			fmt.Println(i, j, s[j:i], dp[j], dict[s[j:i]])
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
