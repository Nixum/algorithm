package array

// 不使用额外空间，且时间复杂度为O(n)的情况
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// 最终一定会相等，如果nums[i]出现在1~n范围内的话
		for nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] - 1 != i {
			res = append(res, i+1)
		}
	}
	return res
}

// 简单写法
func findDisappearedNumbers2(nums []int) []int {
	tmp := make([]int, len(nums)+1)
	for _, n := range nums {
		tmp[n] = 1
	}
	res := make([]int, 0)
	for i := 1; i < len(tmp); i++ {
		if tmp[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
