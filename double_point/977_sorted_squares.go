package double_point

import "sort"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

func sortedSquares2(nums []int) []int {
	res := make([]int, len(nums))
	start := 0
	end := len(nums) - 1
	n := len(nums) - 1
	for start <= end {
		if nums[start] * nums[start] > nums[end] * nums[end] {
			res[n] = nums[start] * nums[start]
			start++
		} else {
			res[n] = nums[end] * nums[end]
			end--
		}
		n--
	}
	return res
}
