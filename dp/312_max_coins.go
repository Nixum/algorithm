package dp

import "algorithm/common"

// 以i,x,x,x,k,x,j为例，
// 假设i,j不可被打破，k是最后一个打破的，此时能得到的值是 i*k*j
// 此时k左边是 i,x,x,x,k , 右边是 k,x,j
// 递归处理，可得递归公式是
// dp[i][j] 表示 范围在 i,j 开区间能得到的硬币数
// dp[i][j] = dp[i][k] + nums[i]*nums[k]*nums[j] + dp[k][j]
// 所以需要保证计算dp[i][j]时，dp[i][k]和dp[k][j]已经被计算出来了（其中i < k < j）
// 确定遍历顺序，根据递推公式画图即可得到，从左到右，从下到上
func maxCoins(nums []int) int {
	dp := make([][]int, len(nums) + 2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums) + 2)
	}
	tmp := make([]int, len(nums) + 2)
	tmp[0] = 1
	tmp[len(tmp)-1] = 1
	for i := 1; i < len(tmp)-1; i++ {
		tmp[i] = nums[i-1]
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		for j := 0; j < len(tmp); j++ {
			for k := i + 1; k < j; k++ {
				left := dp[i][k]
				right := dp[k][j]
				dp[i][j] = common.Max(dp[i][j], left + tmp[i]*tmp[k]*tmp[j] + right)
			}
		}
	}
	return dp[0][len(dp)-1]
}
