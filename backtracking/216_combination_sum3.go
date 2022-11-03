package backtracking

var resInCombinationSum3 [][]int
var pathInCombinationSum3 []int
func combinationSum3(k int, n int) [][]int {
	resInCombinationSum3 = make([][]int, 0)
	pathInCombinationSum3 = make([]int, 0)
	backTrackingInCombinationSum3(k, n, 1)
	return resInCombinationSum3
}

func backTrackingInCombinationSum3(k int, n int, start int) {
	if k == 0 && n == 0 {
		res := make([]int, 0)
		res = append(res, pathInCombinationSum3...)
		resInCombinationSum3 = append(resInCombinationSum3, res)
		return
	}
	for i := start; i <= 9; i++ {
		if n - i < 0 {
			return
		}
		pathInCombinationSum3 = append(pathInCombinationSum3, i)
		backTrackingInCombinationSum3(k - 1, n - i, i + 1)
		pathInCombinationSum3 = pathInCombinationSum3[: len(pathInCombinationSum3) - 1]
	}
	return
}
