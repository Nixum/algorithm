package backtracking

import "sort"

func combinationSum(arr []int, t int) [][]int {
	sort.Ints(arr)
	resInCombinationSum = make([][]int, 0)
	pathInCombinationSum = make([]int, 0)
	backTrackingInCombinationSum(arr, t, 0)
	return resInCombinationSum
}

var resInCombinationSum [][]int
var pathInCombinationSum []int
var sumInCombinationSum int
func backTrackingInCombinationSum(arr []int, t int, start int) {
	if sumInCombinationSum == t {
		res := make([]int, 0)
		res = append(res, pathInCombinationSum...)
		resInCombinationSum = append(resInCombinationSum, res)
		return
	} else if sumInCombinationSum > t {
		return
	}
	// 排序后才能剪枝
	for i := start; i < len(arr) && sumInCombinationSum + arr[i] <= t; i++ {
		sumInCombinationSum += arr[i]
		pathInCombinationSum = append(pathInCombinationSum, arr[i])
		backTrackingInCombinationSum(arr, t, i)
		sumInCombinationSum -= arr[i]
		pathInCombinationSum = pathInCombinationSum[:len(pathInCombinationSum) - 1]
	}
}