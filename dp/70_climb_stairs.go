package dp

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}

// 完全背包解法，求排列，1、2 步 和 2、1 步都是上三个台阶，但是这两种方法是不一样的
// dp[i]：爬到有i个台阶的楼顶，有dp[i]种方法
func climStairs2(n int) int {
	m := 2 // 最多可爬m级台阶
	dp := make([]int, n + 1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; i++ {
			if i - j >= 0 {
				dp[i] += dp[i - j]
			}
		}
	}
	return dp[n]
}