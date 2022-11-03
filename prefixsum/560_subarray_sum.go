package prefixsum

// 统计连续元素的和=k的子数组个数
func subarraySum(nums []int, k int) int {
	// key为前缀和，value为次数
	m := make(map[int]int, 0)
	// 关键！base case
	m[0] = 1
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		// 前缀和
		sum += nums[i]
		// 注意是前缀和 - k
		// 因为是连续元素，即某前缀和 + k = 当前前缀和
		remain := sum - k
		if n, exist := m[remain]; exist {
			res += n
		}
		// 统计前缀和出现的次数
		m[sum] = m[sum] + 1
	}
	return res
}

func subarraySum2(nums []int, k int) int {
	preSum := make([]int, len(nums) + 1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i - 1] + nums[i - 1]
	}
	res := 0
	for i := 0; i < len(preSum); i++ {
		for j := i + 1; j < len(preSum); j++ {
			if preSum[j] - preSum[i] == k {
				res++
			}
		}
	}
	return res
}
