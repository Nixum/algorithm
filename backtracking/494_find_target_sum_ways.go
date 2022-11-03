package backtracking

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	resInFindTargetSumWays = 0
	backTrackingInFindTargetSumWays(nums, target, 0)
	return resInFindTargetSumWays
}

var resInFindTargetSumWays int
func backTrackingInFindTargetSumWays(nums []int, target int, start int) {
	if start == len(nums) {
		if target == 0 {
			resInFindTargetSumWays++
		}
		return
	}
	backTrackingInFindTargetSumWays(nums, target - nums[start], start + 1)
	backTrackingInFindTargetSumWays(nums, target + nums[start], start + 1)
}

// 备忘录消除子问题
func findTargetSumWays2(nums []int, target int) int {
	memoInFindTargetSumWays2 = make(map[string]int, 0)
	return backTrackingInFindTargetSumWays2(nums, target, 0)
}

var memoInFindTargetSumWays2 map[string]int
func backTrackingInFindTargetSumWays2(nums []int, target int, start int) int {
	if start == len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}
	key := fmt.Sprintf("%d-%d", target, start)
	if res, exist := memoInFindTargetSumWays2[key]; exist {
		return res
	}
	res := backTrackingInFindTargetSumWays2(nums, target - nums[start], start + 1) +
		backTrackingInFindTargetSumWays2(nums, target + nums[start], start + 1)
	memoInFindTargetSumWays2[key] = res
	return res
}
