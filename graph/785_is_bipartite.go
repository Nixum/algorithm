package graph

// 二分图：图中有两种颜色，任意两个点间的颜色不同
func isBipartite(graph [][]int) bool {
	visitedInIsBipartite = make([]bool, len(graph))
	colorInIsBipartite = make([]bool, len(graph))
	resInIsBipartite = true
	// 因为图不一定是联通的，所以每个点都需要作为起点遍历一次
	for i := 0; i < len(graph); i++ {
		if !visitedInIsBipartite[i] {
			traverseIsBipartite(graph, i)
		}
	}
	return resInIsBipartite
}

var visitedInIsBipartite []bool
var colorInIsBipartite []bool
var resInIsBipartite bool

func traverseIsBipartite(graph [][]int, v int) {
	if !resInIsBipartite {
		return
	}
	visitedInIsBipartite[v] = true
	for _, nextV := range graph[v] {
		if !visitedInIsBipartite[nextV] {
			// 染色
			colorInIsBipartite[nextV] = !colorInIsBipartite[v]
			// 下一个节点
			traverseIsBipartite(graph, nextV)
		} else {
			// 判断已被访问过的，相邻的节点的颜色是否相同
			if colorInIsBipartite[v] == colorInIsBipartite[nextV] {
				resInIsBipartite = false
				return
			}
		}
	}
}