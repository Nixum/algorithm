package interval_problem

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	// 起点升序，终点降序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	remove := 0
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		a1 := intervals[i][0]
		a2 := intervals[i][1]
		// 区间覆盖
		if left <= a1 && right >= a2 {
			remove++
		}
		// 区间合并
		if left < a1 && right < a2 {
			right = a2
		}
		// 区间重置
		if a1 > right {
			left = a1
			right = a2
		}
	}
	return len(intervals) - remove
}

// 这种思路，在解决大额包含的时候就有问题
// 比如：{3, 4}, {2, 6}, {4, 7}, {1, 8}, {3, 9}, 结果要等于2才对
// 下面这个算法只会在 {3, 4}, {2, 6} 和 {4, 7}, {1, 8} 做移除，此时只删除2个，但实际上{1, 8} 是要删除前3个的
// 解决的方法是需要对区间做合并的操作
func removeCoveredIntervals2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	remove := 0
	for i := 1; i < len(intervals); i++ {
		b1 := intervals[i][0]
		b2 := intervals[i][1]
		a1 := intervals[i - 1][0]
		a2 := intervals[i - 1][1]
		if a1 == b1 && a2 <= b2 {
			fmt.Print(a1, a2, b1, b2, "\n")
			remove++
		} else if a1 <= b1 && a2 == b2 {
			fmt.Print(a1, a2, b1, b2, "\n")
			remove++
		} else if a1 > b1 && a2 < b2 {
			fmt.Print(a1, a2, b1, b2, "\n")
			remove++
		}
	}
	return len(intervals) - remove
}
