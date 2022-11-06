package dp

import "algorithm/common"

// dp[i] 是i之前以nums[i]结尾的最长上升子序列的长度
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = common.Max(dp[i], dp[j] + 1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
