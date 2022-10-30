package dp

// 遇到障碍物，使得dp值为0即可，表示无法走到
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	// 走行和列都是只有一种方式，注意遇到障碍物后，后续都是0了
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		} else {
			break
		}
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] != 1 {
			dp[0][j] = 1
		} else {
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
		}
	}
	return dp[m-1][n-1]
}
