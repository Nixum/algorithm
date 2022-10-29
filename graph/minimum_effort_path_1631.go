package graph

import (
	"algorithm/common"
	"container/heap"
	"math"
)

// 二维数组，从 (0, 0) 到 (n-1, m-1) 的最短路径
func minimumEffortPath(heights [][]int) int {
	effortTo := make([][]int, len(heights))
	for i := 0; i < len(heights); i++ {
		effortTo[i] = make([]int, len(heights[i]))
		for j := 0; j < len(heights[i]); j++ {
			effortTo[i][j] = math.MaxInt64
		}
	}
	direct := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	effortTo[0][0] = 0
	pq := &EffortPathState{}
	heap.Push(pq, []int{0, 0, 0})
	for pq.Len() > 0 {
		cur := heap.Pop(pq).([]int)
		x := cur[0]
		y := cur[1]
		h := cur[2]
		if x == len(heights) - 1 && y == len(heights[0]) - 1 {
			return h
		}
		if h > effortTo[x][y] {
			continue
		}
		edges := adj(heights, x, y, direct)
		for _, edge := range edges {
			nextX := edge[0]
			nextY := edge[1]
			nextEffort := common.Max(effortTo[x][y], common.Abs(heights[x][y] - heights[nextX][nextY]))
			if effortTo[nextX][nextY] > nextEffort {
				effortTo[nextX][nextY] = nextEffort
				heap.Push(pq, []int{nextX, nextY, nextEffort})
			}
		}
	}
	return -1
}

func adj(heights [][]int, x, y int, direct [][]int) [][]int {
	edges := make([][]int, 0)
	for _, d := range direct {
		nextX := d[0] + x
		nextY := d[1] + y
		if nextX < 0 || nextX >= len(heights) || nextY < 0 || nextY >= len(heights[0]) {
			continue
		}
		edges = append(edges, []int{nextX, nextY})
	}
	return edges
}

// ---------

type EffortPathState [][]int

func (h EffortPathState) Len() int           { return len(h) }
func (h EffortPathState) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h EffortPathState) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EffortPathState) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *EffortPathState) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


