package graph

import "container/heap"

// 最小生成树
// 原理：遍历每一个点所关联的边，取权重最小的边
// 拿到该边的下一个点，继续取该点的关联边权重最小的...
// 直到所有点都取完
type Prim struct {
	// 从小到大的优先级队列
	pq *EdgeInPrim
	// 记录哪些节点已经被使用，成为最小生成树的一部分
	inMst []bool
	// 最小权重和
	weightSum int
	// i, j, 3 表示：节点i的所有相邻边数组，有j条边，[3]int表示 from、to、weight
	// i 等于 from
	graph [][][]int
}

func initPrim(graph [][][]int) Prim {
	n := len(graph)
	pq := &EdgeInPrim{}
	weightSum := 0
	inMst := make([]bool, n)
	// 从节点0开始切分
	inMst[0] = true
	cut(graph, inMst, pq, 0)
	// 不断进行切分，向最小生成树中添加边
	for pq.Len() > 0 {
		edges := heap.Pop(pq).([]int)
		to := edges[1]
		weight := edges[2]
		if inMst[to] {
			// 节点to已经添加过了
			continue
		}
		// 边加入最小生成树
		weightSum += weight
		inMst[to] = true
		// 进行新一轮的切分，产生更多横切边
		cut(graph, inMst, pq, to)
	}
	return Prim{
		pq:        pq,
		inMst:     inMst,
		weightSum: weightSum,
		graph:     graph,
	}
}

func cut(graph [][][]int, inMst []bool, pq *EdgeInPrim, s int) {
	for _, edges := range graph[s] {
		to := edges[1]
		if inMst[to] {
			continue
		}
		heap.Push(pq, edges)
	}
}

func (p *Prim) Weight() int {
	return p.weightSum
}

func (p *Prim) AllConnected() bool {
	for i := 0; i < len(p.inMst); i++ {
		if !p.inMst[i] {
			return false
		}
	}
	return true
}

// 优先级队列 ----------------

type EdgeInPrim [][]int

func (h EdgeInPrim) Len() int           { return len(h) }
func (h EdgeInPrim) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h EdgeInPrim) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeInPrim) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *EdgeInPrim) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

