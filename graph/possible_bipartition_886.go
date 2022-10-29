package graph

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n + 1)
	for i := 0; i < len(dislikes); i++ {
		// 注意是双向的
		graph[dislikes[i][0]] = append(graph[dislikes[i][0]], dislikes[i][1])
		graph[dislikes[i][1]] = append(graph[dislikes[i][1]], dislikes[i][0])
	}
	colorInPossibleBipartition = make([]bool, n + 1)
	visitedInPossibleBipartition = make([]bool, n + 1)
	okInPossibleBipartition = true
	for i := 1; i <= n; i++ {
		traversalInPossibleBipartition(graph, i)
	}
	return okInPossibleBipartition
}

var visitedInPossibleBipartition []bool
var colorInPossibleBipartition []bool
var okInPossibleBipartition bool

func traversalInPossibleBipartition(graph [][]int, start int) {
	if !okInPossibleBipartition {
		return
	}
	visitedInPossibleBipartition[start] = true
	for i := 0; i < len(graph[start]); i++ {
		nextV := graph[start][i]
		if !visitedInPossibleBipartition[nextV] {
			colorInPossibleBipartition[nextV] = !colorInPossibleBipartition[start]
			traversalInPossibleBipartition(graph, nextV)
		} else {
			if colorInPossibleBipartition[start] == colorInPossibleBipartition[nextV] {
				okInPossibleBipartition = false
				return
			}
		}
	}
}

