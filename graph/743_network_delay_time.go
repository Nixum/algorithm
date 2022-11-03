package graph

import (
	"algorithm/common"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		graph[i] = make([][]int, 0)
	}
	for _, edges := range times {
		from := edges[0]
		to := edges[1]
		weight := edges[2]
		graph[from] = append(graph[from], []int{to, weight})
	}
	distTo := dijkstra(graph, k)
	// 找出权重最大的最短路径
	res := 0
	for i := 1; i < len(distTo); i++ {
		if distTo[i] == math.MaxInt64 {
			return -1
		}
		res = common.Max(res, distTo[i])
	}
	return res
}
