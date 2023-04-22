package array

// 给定一个正整数数组 arr，求 arr[i] - arr[j] 的最大值，
// 其中 i < j
func getMaxDiff(nums []int) int {
	max := 0
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] - nums[i] > max {
			max = nums[j] - nums[i]
		}
		if nums[j] < nums[i] {
			i = j
		}
	}
	return max
}
