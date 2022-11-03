package graph

// 有向无环图的逆后序排序就是其拓扑排序
func findOrder(numCourses int, prerequisites [][]int) []int {
	resInFindOrder = make([]int, 0)
	visitedInFindOrder = make([]bool, numCourses)
	pathInFindOrder = make([]bool, numCourses)
	hasCycleInFindOrder = false

	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}
	for i := 0; i < numCourses; i++ {
		traverseInFindOrder(graph, i)
	}
	if hasCycleInFindOrder {
		return []int{}
	}
	for i := 0; i < len(resInFindOrder) / 2; i++ {
		resInFindOrder[i], resInFindOrder[len(resInFindOrder) - i - 1] = resInFindOrder[len(resInFindOrder) - i - 1], resInFindOrder[i]
	}
	return resInFindOrder
}

var resInFindOrder []int
var visitedInFindOrder []bool
var pathInFindOrder []bool
var hasCycleInFindOrder bool

func traverseInFindOrder(graph [][]int, start int) {
	if pathInFindOrder[start] {
		hasCycleInFindOrder = true
	}
	if visitedInFindOrder[start] || hasCycleInFindOrder {
		return
	}
	visitedInFindOrder[start] = true
	pathInFindOrder[start] = true
	for _, nextIndex := range graph[start] {
		traverseInFindOrder(graph, nextIndex)
	}
	resInFindOrder = append(resInFindOrder, start)
	pathInFindOrder[start] = false
}
