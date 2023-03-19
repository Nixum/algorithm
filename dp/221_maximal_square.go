package dp

import (
	"algorithm/common"
	"math"
)

// 二维DP
// dp[i][j]: 以下标[i][j]为右下角能构成的最大正方形的边长
// 此时有两种情况：
// 1. matrix[i][j] == 0
//   此时无法构成正方形，dp[i][j] = 0
// 2. matrix[i][j] == 1
//   要判断包含matrix[i][j]能否加入，
//   就要看dp[i-1][j]、dp[i][j-1]、dp[i-1][j-1]能构成的正方形的大小
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	res := math.MinInt64
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			res = common.Max(res, dp[i][0])
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			res = common.Max(res, dp[0][j])
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = common.Min(common.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				res = common.Max(res, dp[i][j])
			}
		}
	}
	return res * res
}
