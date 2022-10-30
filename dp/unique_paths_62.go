package dp

// 递推公式：dp[i][j]表示从(0,0)到(i,j)总共有dp[i][j]条不同路径
// 因为机器人只能从往右和下前进，
// 所以，dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	// 走行和列都是只有一种方式
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
		}
	}
	return dp[m-1][n-1]
}
