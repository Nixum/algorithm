package backtracking

import "sort"

func combinationSum2(arr []int, t int) [][]int {
	resInCombinationSum2 = make([][]int, 0)
	pathInCombinationSum2 = make([]int, 0)
	sort.Ints(arr)
	flag := make([]bool, len(arr))
	backTrackingInCombinationSum2(arr, t, 0, flag)
	return resInCombinationSum2
}

var resInCombinationSum2 [][]int
var pathInCombinationSum2 []int
func backTrackingInCombinationSum2(arr []int, t int, start int, flag []bool) {
	if t == 0 {
		res := make([]int, 0)
		res = append(res, pathInCombinationSum2...)
		resInCombinationSum2 = append(resInCombinationSum2, res)
		return
	} else if t < 0 {
		return
	}
	for i := start; i < len(arr); i++ {
		// 防止同层重复取，比如【1，1，2】，t=3
		if i - 1 >= 0 && arr[i] == arr[i - 1] && flag[i - 1] == false {
			continue
		}
		flag[i] = true
		pathInCombinationSum2 = append(pathInCombinationSum2, arr[i])
		backTrackingInCombinationSum2(arr, t - arr[i], i + 1, flag)
		flag[i] = false
		pathInCombinationSum2 = pathInCombinationSum2[:len(pathInCombinationSum2) - 1]
	}
}
