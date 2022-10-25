package graph

import "sort"

// 最小生成树
func minimumCost(n int, connections [][]int) int {
	// 节点序号从1开始
	uf := InitUF(n + 1)
	// 对所有边按权重从小到大排列
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	min := 0
	for _, edge := range connections {
		from := edge[0]
		to := edge[1]
		weight := edge[2]
		// 说明产生环
		if uf.Connected(from, to) {
			continue
		}
		min += weight
		uf.Union(from, to)
	}
	// 保证所有节点都被连通
	// 因为要联通所有节点，按理说 uf.count() == 1 说明所有节点被连通
	// 但因为节点 0 没有被使用，所以 0 会额外占用一个连通分量
	if uf.Count() == 2 {
		return min
	}
	return -1
}
