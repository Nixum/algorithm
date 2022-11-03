package dp

import "algorithm/common"

// 递推公式，dp[i]：拆分数字i，可以得到的最大乘积
// 总共有两种渠道得到dp[i]：
// j * (i - j) 和 dp[i - j] * j
// 一个表示 拆成两个数，另一个表示拆成3个数以上
func integerBreak(n int) int {
	dp := make([]int, n + 1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i - 1; j++ {
			//dp[i] = common.Max(dp[i], dp[i - j] * dp[j])
			dp[i] = common.Max(dp[i], common.Max((i - j) * j, dp[i - j] * j))
		}
	}
	return dp[n]
}
