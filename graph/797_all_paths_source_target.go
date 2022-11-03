package graph

// 无环图不需要visited，有向含环图才需要visited
func allPathsSourceTarget(graph [][]int) [][]int {
	resInAllPathsSourceTarget = make([][]int, 0)
	pathInAllPathsSourceTarget = make([]int, 0)
	trackingInAllPathsSourceTarget(graph, 0)
	return resInAllPathsSourceTarget
}

var resInAllPathsSourceTarget [][]int
var pathInAllPathsSourceTarget []int
func trackingInAllPathsSourceTarget(graph [][]int, start int) {
	// 终点是固定的，所以走到终点就得加结果
	if start == len(graph) - 1 {
		res := make([]int, 0)
		res = append(res, pathInAllPathsSourceTarget...)
		res = append(res, start)
		resInAllPathsSourceTarget = append(resInAllPathsSourceTarget, res)
		return
	}
	// 和回溯有点像，差异就是图关注的是节点，回溯关注的是树枝
	pathInAllPathsSourceTarget = append(pathInAllPathsSourceTarget, start)
	for i := 0; i < len(graph[start]); i++ {
		trackingInAllPathsSourceTarget(graph, graph[start][i])
	}
	pathInAllPathsSourceTarget = pathInAllPathsSourceTarget[:len(pathInAllPathsSourceTarget) - 1]
}