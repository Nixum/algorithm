package backtracking

var resInCombine [][]int
var itemInCombine []int
func combine(n int, k int) [][]int {
	resInCombine = make([][]int, 0)
	itemInCombine = make([]int, 0)
	backTrackingInCombine(n, k, 1)
	return resInCombine
}

func backTrackingInCombine(n int, k int, start int) {
	if k == 0 {
		res := make([]int, 0)
		res = append(res, itemInCombine...)
		resInCombine = append(resInCombine, res)
		return
	}
	for i := start; i <= n; i++ {
		itemInCombine = append(itemInCombine, i)
		backTrackingInCombine(n, k - 1, i + 1)
		itemInCombine = itemInCombine[:len(itemInCombine) - 1]
	}
	return
}