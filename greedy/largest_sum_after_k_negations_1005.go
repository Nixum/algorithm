package greedy

import (
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		t1 := nums[i]
		if t1 < 0 {
			t1 = -nums[i]
		}
		t2 := nums[j]
		if t2 > 0 {
			t2 = nums[j]
		}
		// 必须从大排到小，这样反转负数才能先反转大的
		return t1 > t2
	})
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] = -nums[i]
		}
	}
	if k % 2 == 1 {
		nums[0] = -nums[0]
	}
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
