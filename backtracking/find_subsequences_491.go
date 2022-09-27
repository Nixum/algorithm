package backtracking

func findSubsequences(nums []int) [][]int {
	resInFindSubsequences = make([][]int, 0)
	pathInFindSubsequences = make([]int, 0)
	backTrackingInInFindSubsequences(nums, 0)
	return resInFindSubsequences
}

var resInFindSubsequences [][]int
var pathInFindSubsequences []int
func backTrackingInInFindSubsequences(nums []int, start int) {
	if len(pathInFindSubsequences) >= 2 {
		res := make([]int, 0)
		res = append(res, pathInFindSubsequences...)
		resInFindSubsequences = append(resInFindSubsequences, res)
	}
	if start >= len(nums) {
		return
	}
	// 注意是同层级相同数字就得去重
	isUsed := make(map[int]int, len(nums))
	for i := start; i < len(nums); i++ {
		if _, exist := isUsed[nums[i]]; exist {
			continue
		}
		if len(pathInFindSubsequences) > 0 && pathInFindSubsequences[len(pathInFindSubsequences) - 1] > nums[i] {
			continue
		}
		pathInFindSubsequences = append(pathInFindSubsequences, nums[i])
		isUsed[nums[i]] = 1
		backTrackingInInFindSubsequences(nums, i + 1)
		pathInFindSubsequences = pathInFindSubsequences[:len(pathInFindSubsequences) - 1]
	}
}