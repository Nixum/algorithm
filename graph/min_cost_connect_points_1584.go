package graph

import (
	"algorithm/common"
	"sort"
)

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	edges := make([][]int, 0)
	// 构造出所有边
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges,
				[]int{i, j, common.Abs(points[i][0] - points[j][0]) + common.Abs(points[i][1] - points[j][1])})
		}
	}
	// 边从小到大排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	res := 0
	uf := InitUF(n)
	// 生成最小生成树
	for i := 0; i < len(edges); i++ {
		if uf.Connected(edges[i][0], edges[i][1]) {
			continue
		}
		uf.Union(edges[i][0], edges[i][1])
		res += edges[i][2]
	}
	return res
}
