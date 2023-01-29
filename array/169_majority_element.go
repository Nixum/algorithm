package array

import "sort"

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。
// 多数元素是指在数组中出现次数 大于 n/2 的元素。
// 要求时间复杂度为 O(n)、空间复杂度为 O(1)
func majorityElement(nums []int) int {
	sort.Ints(nums)
	numsLen := len(nums)
	res := nums[0]
	resN := 1
	for i := 1; i < numsLen; i++ {
		if nums[i] == res {
			resN++
		} else {
			res = nums[i]
			resN = 1
		}
		if resN > numsLen / 2 {
			break
		}
	}
	return res
}

// 另一种解法，摩尔投票，需要确保数组中必定有数量大于n/2的数
// 因此 这个数的个数 - 其余元素个数的总和，结果必定 >= 1
// 将第一个元素设为标志值，数量为1
// 遍历元素，当元素与标志值不同时，数量--，相同则++
// 当数量为0时，更换标志值，数量值为1
// 最后返回的标志的数量就会大于 n/2
func majorityElement2(nums []int) int {
	n := nums[0]
	nLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == n {
			nLen++
		} else {
			nLen--
			if nLen == 0 {
				n = nums[i]
				nLen = 1
			}
		}
	}
	return n
}