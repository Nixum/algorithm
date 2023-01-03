package dp

import "algorithm/common"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dpMax := nums[0]
	dpMin := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dpMax = common.Max(common.Max(dpMax * nums[i], dpMin * nums[i]), nums[i])
		dpMin = common.Min(common.Min(dpMin * nums[i], dpMax * nums[i]), nums[i])
		res = common.Max(res, dpMax)
	}
	return res
}

// 其他思路：
// 所要求的乘积最大，就要尽可能多的数字相乘，没有负数，那就直接所有数字相乘
// 如果有负数，就分情况，一种是奇数个负数，一种是偶数个负数，如果是偶数个负数，直接相乘即可，
// 如果是奇数个负数，因为要尽可能多的数字相乘，则有两种情况取最大值，
// 一种是第一个负数之后的数字相乘得到的值，一种是倒数第一个负数之前的数字相乘得到的值
// 取两者最大即可，然后相乘过程中，还需要主要遇到0的场景
func maxProduct1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	// 正序累乘
	max := 1
	for i := 0; i < len(nums); i++ {
		// 一直累乘，如果负数是奇数个，那乘于最后一个负数之后，再继续乘都不会得到最大值了
		max *= nums[i]
		res = common.Max(max, res)
		if max == 0 {
			max = 1
		}
	}
	// 倒序累乘
	max = 1
	for i := len(nums) - 1; i >= 0; i-- {
		// 一直累乘，如果负数是奇数个，那乘于第一个负数之后，再继续乘都不会得到最大值了
		max *= nums[i]
		res = common.Max(max, res)
		if max == 0 {
			max = 1
		}
	}
	return res
}

// 错误的思路
// dp[i][j]: 下标i到j的最大乘积
// dp[i][j] = max(dp[i][j], dp[i][j - 1] * nums[j], dp[i - 1][j] * nums[i])
func maxProduct2(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	dp[0][0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[0][i] = dp[0][i - 1] * nums[i]
	}
	for i := 1; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = common.Max(dp[i][j], common.Max(dp[i][j - 1] * nums[j], dp[i - 1][j] * nums[i]))
			//dp[i][j] = common.Max(dp[i][j], dp[i][j - 1] * nums[j])
		}
	}
	return dp[len(nums) - 1][len(nums) - 1]
}
