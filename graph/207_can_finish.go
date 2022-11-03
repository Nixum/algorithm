package graph

// numCourses是课程数，即节点的数量
// 当出现环的时候，结果为false，所以这道题的本质是判断有没有环
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}
	hasCycleInCanFinish = false
	visitedInCanFinish = make([]bool, numCourses)
	pathInCanFinish = make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		// 每门课程都判断一下
		traverseInCanFinish(graph, i)
	}
	return !hasCycleInCanFinish
}

var hasCycleInCanFinish bool
var visitedInCanFinish []bool // 有向有环图需要visited记录来判断节点是否走过
var pathInCanFinish []bool // 记录每个节点是否遍历完成，说明所有课程都修完了
func traverseInCanFinish(graph [][]int, start int) {
	if pathInCanFinish[start] {
		hasCycleInCanFinish = true
	}
	if visitedInCanFinish[start] || hasCycleInCanFinish {
		return
	}
	visitedInCanFinish[start] = true
	pathInCanFinish[start] = true
	for i := 0; i < len(graph[start]); i++ {
		traverseInCanFinish(graph, graph[start][i])
	}
	pathInCanFinish[start] = false
}