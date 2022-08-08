package interval_problem

import (
	"algorithm/common"
	"sort"
)

// 1. 排序，注意起点和终点的顺序，比如当起点相等时，终点是要正序还是倒序
// 2. 分情况去分析
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		a1 := intervals[i][0]
		a2 := intervals[i][1]
		lastRes := res[len(res) - 1]
		if a1 <= lastRes[1] {
			res[len(res) - 1][1] = common.Max(a2, lastRes[1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
