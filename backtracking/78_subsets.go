package backtracking

func subsets(nums []int) [][]int {
	resInSubsets = make([][]int, 0)
	pathInSubsets = make([]int, 0)
	backTrackingInSubsets(nums, 0)
	return resInSubsets
}

var resInSubsets [][]int
var pathInSubsets []int
func backTrackingInSubsets(nums []int, start int) {
	res := make([]int, 0)
	res = append(res, pathInSubsets...)
	resInSubsets = append(resInSubsets, res)
	if start >= len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		pathInSubsets = append(pathInSubsets, nums[i])
		backTrackingInSubsets(nums, i + 1)
		pathInSubsets = pathInSubsets[:len(pathInSubsets) - 1]
	}
}
