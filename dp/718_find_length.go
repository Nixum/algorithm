package dp

// 因为边界问题，所以由i-1,j-1比较
// dp[i][j] ：以下标i - 1为结尾的A，和以下标j - 1为结尾的B，最长重复子数组长度为dp[i][j]
// 因为要求 连续 子序列相同，所以可以由前面一个dp推出
// 递推公式：
// dp[i][j] = dp[i - 1][j - 1] + 1
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1) + 1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2) + 1)
	}
	res := 0
	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i - 1] == nums2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
