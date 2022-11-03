package backtracking

func permute(nums []int) [][]int {
	resInPermute = make([][]int, 0)
	pathInPermute = make([]int, 0)
	isUsed := make([]bool, len(nums))
	backTrackingInPermute(nums, isUsed)
	return resInPermute
}

var resInPermute [][]int
var pathInPermute []int
func backTrackingInPermute(nums []int, isUsed []bool) {
	if len(pathInPermute) == len(nums) {
		res := make([]int, 0)
		res = append(res, pathInPermute...)
		resInPermute = append(resInPermute, res)
		return
	}
	for i := 0; i < len(nums); i++ {
		if isUsed[i] {
			continue
		}
		isUsed[i] = true
		pathInPermute = append(pathInPermute, nums[i])
		backTrackingInPermute(nums, isUsed)
		isUsed[i] = false
		pathInPermute = pathInPermute[:len(pathInPermute) - 1]
	}
}
