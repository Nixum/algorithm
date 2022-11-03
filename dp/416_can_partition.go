package dp

import "algorithm/common"

// 递推公式：
// dp[j]的含义：最大可凑成i的子集和，索引表示目标和，
// 所以，dp[j] = max(dp[j], dp[j - nums[i]] + nums[i])
// max(最大可凑成i的子集和, 目标和 - 当前数 的最大可凑成的子集和 + 当前数)
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum % 2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target + 1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = common.Max(dp[j], dp[j - nums[i]] + nums[i])
		}
	}
	return dp[target] == target
}

// 能AC，但会超时
func canPartition1(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum % 2 != 0 {
		return false
	}
	target := sum / 2
	sumInCanPartition = 0
	resInCanPartition = false
	backTrackingInCanPartition(nums, target, 0)
	return resInCanPartition
}

var sumInCanPartition int
var resInCanPartition bool
func backTrackingInCanPartition(nums []int, t int, start int) {
	if start >= len(nums) || resInCanPartition || sumInCanPartition > t{
		return
	}
	if sumInCanPartition == t {
		resInCanPartition = true
		return
	}
	for i := start; i < len(nums); i++ {
		if sumInCanPartition + nums[i] > t {
			continue
		}
		sumInCanPartition += nums[i]
		if sumInCanPartition == t {
			resInCanPartition = true
			return
		}
		backTrackingInCanPartition(nums, t, i + 1)
		sumInCanPartition -= nums[i]
	}
}
