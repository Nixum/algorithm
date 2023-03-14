package double_point

import "sort"

// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，
// 如果对这个子数组进行升序排序，那么整个数组都会变为升序排序，返回这个子数组的长度
// 思路：
// 将数组分成三段，第一、三段肯定是升序的，目标是找出第二段的起始位置
// 第一段的end，从左到右遍历，其值肯定比前面所有数都大，必定会走到第三段的开始
// end就是最后一个小于max的元素的位置
// 第三段的begin，从右到左遍历，其值肯定比后面所有数都小，必定会走到第一段的结尾
// begin就是最后一个大于min的元素的位置
func findUnsortedSubarray(nums []int) int {
	numLen := len(nums)
	max := nums[0]
	min := nums[numLen-1]
	// 这两个值初始化不能写错
	end := -1
	begin := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}

		if nums[numLen-i-1] > min {
			begin = numLen-i-1
		} else {
			min = nums[numLen-i-1]
		}
	}
	return end - begin + 1
}

func findUnsortedSubarray2(nums []int) int {
	ns := make([]int, 0, len(nums))
	ns = append(ns, nums...)
	sort.Ints(ns)
	begin := 0
	end := len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i] == ns[i] {
			begin = i
		} else {
			break
		}
	}
	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] == ns[i] {
			end = i
		} else {
			break
		}
	}
	if end <= begin {
		return 0
	}
	return end - begin + 1
}