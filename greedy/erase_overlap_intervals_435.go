package greedy

import "sort"

// 思路：总区间减去非重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	notOverlap := 1
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			notOverlap++
		}
	}
	return len(intervals) - notOverlap
}

// 错误的思路，想得太简单了
func eraseOverlapIntervals2(intervals [][]int) int {
	overlap := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i - 1][1] {
			overlap++
		}
	}
	return len(intervals) - 1 - overlap
}