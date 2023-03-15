package prefixsum

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
// 不使用除法，且在 O(n) 时间复杂度内完成，这种空间复杂度是 O(n)
// 思路，计算nums[i]的前缀积和后缀积
// 不包含nums[i]本身的积就是 preSum[i-1] * postSum[i+1]
func productExceptSelf(nums []int) []int {
	prefixMulti := make([]int, len(nums))
	postfixMulti := make([]int, len(nums))
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
		prefixMulti[i] = sum
	}
	sum = 1
	for i := len(nums)-1; i >= 0; i-- {
		sum *= nums[i]
		postfixMulti[i] = sum
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = postfixMulti[i+1]
		} else if i == len(nums)-1 {
			res[i] = prefixMulti[i-1]
		} else {
			res[i] = prefixMulti[i-1] * postfixMulti[i+1]
		}
	}
	return res
}

// 不使用除法，且在 O(n) 时间复杂度内完成，
// 且在不算结果数组，空间复杂度是 O(1)
// 计算 不包含 nums[i] 本身的前缀和和后缀和
func productExceptSelf2(nums []int) []int {
	// res数组不算的情况下，不使用额外空间，其实就是用res数组来装了
	res := make([]int, len(nums))
	sum := 1
	for i := 1; i <= len(nums); i++ {
		res[i-1] = sum
		sum *= nums[i-1]
	}
	sum = 1
	for i := len(nums); i >= 1; i-- {
		res[i-1] *= sum
		sum *= nums[i-1]
	}
	return res
}