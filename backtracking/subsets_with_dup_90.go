package backtracking

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	resInSubsetsWithDup = make([][]int, 0)
	pathInSubsetsWithDup = make([]int, 0)
	sort.Ints(nums)
	backTrackingInSubsetsWithDup(nums, 0, false)
	return resInSubsetsWithDup
}

var resInSubsetsWithDup [][]int
var pathInSubsetsWithDup []int
func backTrackingInSubsetsWithDup(nums []int, start int, isUsed bool) {
	res := make([]int, 0)
	res = append(res, pathInSubsetsWithDup...)
	resInSubsetsWithDup = append(resInSubsetsWithDup, res)
	if start >= len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		if i - 1 >= 0 && nums[i - 1] == nums[i] && !isUsed {
			continue
		}
		isUsed = true
		pathInSubsetsWithDup = append(pathInSubsetsWithDup, nums[i])
		backTrackingInSubsetsWithDup(nums, i + 1, isUsed)
		pathInSubsetsWithDup = pathInSubsetsWithDup[:len(pathInSubsetsWithDup) - 1]
		isUsed = false
	}
}