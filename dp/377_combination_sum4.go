package dp

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target + 1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i - nums[i] >= 0 {
				dp[i] += dp[i - nums[i]]
			}
		}
	}
	return dp[target]
}