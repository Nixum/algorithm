package array

import "algorithm/common"

// 求数组中最长连续序列的长度，不要求连续
// 要求时间复杂度是O(n)
func longestConsecutive(nums []int) int {
	// 因为是连续的序列，所以可以判断 当前值-1 是否是最小的，如果不是，则无需统计
	// 如果是，才需要统计，以此减少遍历次数
	set := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}
	max := 0
	for _, n := range nums {
		// 因为这里每个元素只判断了一次，所以时间复杂度是O(n)
		if !set[n-1] {
			count := 1
			for set[n] {
				count++
				n++
			}
			max = common.Max(max, count)
		}
	}
	return max
}

// 思路是对的，但是超时了，时间复杂度是O（n^2）
func longestConsecutive1(nums []int) int {
	set := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		base := nums[i]
		sum := 1
		for j := 1; j <= len(nums); j++ {
			base++
			if set[base] {
				sum++
			} else {
				break
			}
		}
		max = common.Max(max, sum)
	}
	return max
}
