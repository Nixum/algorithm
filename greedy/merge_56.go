package greedy

import (
	"algorithm/common"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			// 这个很关键
			right = common.Max(intervals[i][1], right)
		} else {
			res = append(res, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	res = append(res, []int{left, right})
	return res
}