package graph

import (
	"container/heap"
	"math"
)

// 需要准备的东西：
// 1. 存放起点到每个点的最短距离数组
// 2. 存放到下一个节点的最短距离的 优先级队列
// 3. 每个点到其周围点的边的数组
// 思路：
// 从优先级队列里获取下一个要处理的节点，判断上一个点到当前点的最短距离是否比结果的大
// 如果大，则从优先级队列里取下一个要处理的点
// 如果小，说明可以走这个点; 遍历该点关联的所有点，
// 判断下一个点的最短距离 是否比 当前最短距离加下一个点的权重 大
// 如果是，更新下一个点的最短距离，并加入优先级队列
// 循环从优先级队列中获取下一个点

// dijkstra算法，得到起点到每个点的最短距离
// 适用于 加权有向图
func dijkstra(graph [][][]int, start int) []int {
	n := len(graph)
	distTo := make([]int, n)
	for i := 0; i < len(distTo); i++ {
		distTo[i] = math.MaxInt64
	}
	distTo[start] = 0
	pq := &stateInDijkstra{}
	pq.Push(dijkstraState{
		id:            start,
		distFromStart: 0,
	})
	for pq.Len() > 0 {
		curState := heap.Pop(pq).(dijkstraState)
		curId := curState.id
		curDistFromStart := curState.distFromStart
		// 已经有一条更短的路径到达当前节点
		if curDistFromStart > distTo[curId] {
			continue
		}
		// 将当前节点的相邻节点加入队列
		for i := 0; i < len(graph[curId]); i++ {
			// 看看从当前节点到下一个节点的距离是否会更短
			nextId := graph[curId][i][0]
			distToNext := distTo[curId] + graph[curId][i][1]
			if distTo[nextId] > distToNext {
				// 更新dp table
				distTo[nextId] = distToNext
				// 对应的节点和距离入队
				heap.Push(pq, dijkstraState{
					id:            nextId,
					distFromStart: distToNext,
				})
			}
		}
	}
	return distTo
}

type dijkstraState struct {
	// 节点
	id int
	// 从start到当前节点的权重
	distFromStart int
}

// 优先级队列 ----------------

type stateInDijkstra []dijkstraState

func (h stateInDijkstra) Len() int           { return len(h) }
func (h stateInDijkstra) Less(i, j int) bool { return h[i].distFromStart < h[j].distFromStart }
func (h stateInDijkstra) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *stateInDijkstra) Push(x interface{}) {
	*h = append(*h, x.(dijkstraState))
}

func (h *stateInDijkstra) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

