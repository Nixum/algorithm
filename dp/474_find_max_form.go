package dp

import "algorithm/common"

// dp[i][j]：i个0，j个1的strs的最大子集的大小
// 递推公式：dp[i][j] = max(dp[i][j], dp[i-zero]][j-one] + 1)
// 类比成背包，zero和one相当于物品的重量，字符串本身的个数是物品价值，都是为 1
func findMaxForm(strs []string, m int, n int) int {
	zero := make([]int, 0)
	one := make([]int, 0)
	for _, str := range strs {
		zeroNum := 0
		oneNum := 0
		for _, b := range str {
			if b - '0' == 0 {
				zeroNum++
			} else {
				oneNum++
			}
		}
		zero = append(zero, zeroNum)
		one = append(one, oneNum)
	}
	dp := make([][]int, m + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n + 1)
	}
	for i := 0; i < len(strs); i++ {
		for j := m; j >= zero[i]; j-- {
			for k := n; k >= one[i]; k-- {
				dp[j][k] = common.Max(dp[j][k], dp[j - zero[i]][k - one[i]] + 1)
			}
		}
	}
	return dp[m][n]
}
