package dp

// 字符串中回文子串的个数
// dp[i][j], 以i，j左闭右闭区间是否为回文子串
// 分成:
// s[i] == s[j], 有三种情况
// 1: s[i] == s[j] && i == j, 即单个字符，所以dp[i][j] = true
// 2: s[i] == s[j] && i + 1 == j, 即两个字符，所以dp[i][j] = true
// 3: s[i] == s[j] && j - i > 1 && dp[i+1][j-1] == true, 即三个字符以上 dp[i][j] = true
// s[i] != s[j], dp[i][j] = false
// 由于 dp[i][j] = dp[i+1][j-1] 得出，所以遍历顺序是从下到上，从左到右
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j - 1 <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i + 1][j - 1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}
