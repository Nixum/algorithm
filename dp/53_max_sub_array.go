package dp

import "algorithm/common"

// dp[i], 包括下标i之前的最大连续子序列和为dp[i]
// dp[i] = max(nums[i], dp[i - 1] + nums[i])
// 比较 从头开始计算的当前子序列和 和 前一个最大连续子序列和加上当前值
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = common.Max(nums[i], dp[i - 1] + nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// 错误的思路，这样会变成累加，而不会产生中断
// 因为当和为负数时，再加就变少了，需要重新取值
// 如果是下面这种写法，就会跳过该值继续累加
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = common.Max(dp[i - 1] , dp[i - 1] + nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
