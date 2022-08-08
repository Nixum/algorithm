package interval_problem

import "sort"

// 算非交叉区间有多少个即可得到结果
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		// 注意这里是没等于的
		if end < points[i][0] {
			res++
			end = points[i][1]
		}
	}
	return res
}
