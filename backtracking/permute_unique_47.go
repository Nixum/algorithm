package backtracking

import "sort"

func permuteUnique(nums []int) [][]int {
	resInPermuteUnique = make([][]int, 0)
	pathInPermuteUnique = make([]int, 0)
	isUsed := make([]bool, len(nums))
	sort.Ints(nums)
	backTrackingInPermuteUnique(nums, isUsed)
	return resInPermuteUnique
}

var resInPermuteUnique [][]int
var pathInPermuteUnique []int
func backTrackingInPermuteUnique(nums []int, isUsed []bool) {
	if len(pathInPermuteUnique) == len(nums) {
		res := make([]int, 0)
		res = append(res, pathInPermuteUnique...)
		resInPermuteUnique = append(resInPermuteUnique, res)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 树枝的
		if isUsed[i] {
			continue
		}
		// 树层的
		if i > 0 && nums[i] == nums[i - 1] && isUsed[i - 1] {
			continue
		}
		// 这样写也能ac，但是比上面的效率差点
		//if i > 0 && nums[i] == nums[i - 1] && isUsed[i - 1] {
		//	break
		//}
		isUsed[i] = true
		pathInPermuteUnique = append(pathInPermuteUnique, nums[i])
		backTrackingInPermuteUnique(nums, isUsed)
		isUsed[i] = false
		pathInPermuteUnique = pathInPermuteUnique[:len(pathInPermuteUnique) - 1]
	}
}