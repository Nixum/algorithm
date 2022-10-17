package greedy

import (
	"algorithm/common"
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] <= points[j][0]
	})
	fmt.Println(points)
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i - 1][1] {
			fmt.Println()
			res++
		} else {
			fmt.Println(points[i - 1][1], points[i][1])
			points[i][1] = common.Min(points[i - 1][1], points[i][1])
		}
	}
	return res
}

// 错误的思路，难点在于打爆气球之后，即res++，end的位置要重新记录
// 由于上面是直接修改上一个气球的end位置的，当气球打爆时，遍历到了当前的气球，此时end位置重新计算
func findMinArrowShots2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] <= points[j][0]
	})
	end := points[0][1]
	res := 1
	for i := 1; i < len(points); i++ {
		fmt.Println(end)
		//if end > points[i][0] && end <= points[i][1] {
		//	end = common.Min(end, points[i][1])
		//} else if end < points[i][0] {
		//	end = points[i][1]
		//	res++
		//}
		if points[i][0] > end {
			res++
			// 需要加上这句才行
			end = points[i][1]
		} else {
			// 这里有问题，当res++之后，end仍然记录上一次的最小，导致多算一次
			end = common.Min(end, points[i][1])
		}
	}
	return res
}