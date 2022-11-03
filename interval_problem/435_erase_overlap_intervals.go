package interval_problem

import (
	"fmt"
	"sort"
)

// 另一种方案，总区间个数 - 非交叉区间个数 = 移除的区间数
func eraseOverlapIntervals2(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1]
	count := 1 // 非交叉区间个数
	for i := 1; i < len(intervals); i++ {
		// 注意这里是等于的，因为题干里如果边界值相等不算覆盖
		if end <= intervals[i][0] {
			count++
			end = intervals[i][1]
		}
	}
	return len(intervals) - count
}

// 因为是要去除区间，所以按终点升序，然后比较终点即可
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res := 0
	tmp := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		a1 := intervals[i][0]
		a2 := intervals[i][1]
		if a1 < tmp[1] {
			res++
		} else {
			tmp[0] = a1
			tmp[1] = a2
		}
	}
	return res
}

// 错误的思路, 因为是按起点升序，终点升序得，这个算法会导致每次都只删除当前一个，保留前一个
// [-73 -26] [-65 -11] [-63 2] [-62 -49] [-52 31] [-40 -26] [-31 49] [30 47] [58 95] [66 98] [82 97] [95 99]
// 经过该算法后变成：[-73 -26] [30 47] [58 95] [82 97]
// 实际上答案应该是：[-62 -49] [-40 -26] [30 47] [58 95] [82 97]
func eraseOverlapIntervals3(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	tmp := []int{intervals[0][0], intervals[0][1]}
	fmt.Println(intervals)
	for i := 1; i < len(intervals); i++ {
		a1 := intervals[i][0]
		a2 := intervals[i][1]
		if a1 < tmp[1] {
			res++
		} else {
			tmp[0] = a1
			tmp[1] = a2
		}
	}
	return res
}
