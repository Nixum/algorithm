package dp

import "algorithm/common"

// 尽量让石头分成重量相同的两堆
// 递推公式：
// dp[j]：容量为j的背包最多能背的石头重量
// 最后一块石头的重量 = 石头总量 - dp[sum/2] * 2
func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}
	target := sum / 2
	dp := make([]int, target + 1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = common.Max(dp[j], dp[j - stones[i]] + stones[i])
		}
	}
	return sum - dp[len(dp) - 1] * 2
}