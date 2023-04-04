package week2

import "algorithm/common"

// dp[i], 第i个字符包括i的最长平衡子串长度
// 如果 s[i] == '0', dp[i] = 0
// 如果 s[i] == '1' && s[i-dp[i-1]-1] == '0', dp[i] = dp[i-1] + 2
// 如果 s[i] == '1' && s[i-dp[i-1]-1] != '0', dp[i] = 0
func findTheLongestBalancedSubstring(s string) int {
	max := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == '1' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '0' {
			dp[i] = dp[i-1] + 2
		}
		max = common.Max(max, dp[i])
	}
	return max
}

// 空间复杂度为O(1)的解法
func findTheLongestBalancedSubstring1(s string) int {
	max := 0
	// 连续0的个数，连续1的个数
	pre, cur := 0, 0
	for i, c := range s {
		cur++
		if i == len(s)-1 || byte(c) != s[i+1] {
			if c == '1' {
				max = common.Max(max, common.Min(pre, cur) * 2)
			}
			pre = cur
			cur = 0
		}
	}
	return max
}