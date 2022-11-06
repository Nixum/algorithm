package dp

// dp[i] 是i之前以nums[i]结尾的最长上升子序列的长度
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i - 1] {
			dp[i] = dp[i - 1] + 1
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
