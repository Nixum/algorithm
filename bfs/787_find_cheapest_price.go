package bfs

import (
	"algorithm/common"
	"math"
)

// BFS 思路也可以，但是在leetcode会 out of memory, 时间复杂度O(n^k)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][][]int)
	for i := 0; i < len(flights); i++ {
		s := flights[i][0]
		d := flights[i][1]
		p := flights[i][2]
		graph[s] = append(graph[s], []int{d, p})
	}
	step := 0
	queue := make([][]int, 0)
	minPrice := math.MaxInt64
	queue = append(queue, []int{src, 0})
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur, q := QueuePop(queue)
			queue = q
			if cur[0] == dst {
				minPrice = common.Min(minPrice, cur[1])
			}
			if _, exist := graph[cur[0]]; !exist {
				continue
			}
			for _, next := range graph[cur[0]] {
				if cur[1] + next[1] > minPrice {
					continue
				}
				queue = append(queue, []int{next[0], cur[1] + next[1]})
			}
			// 就算加了这句也会 out of memory
			cur = nil
		}
		if step > k {
			break
		}
		// 这句一定在后面, 表示算完了当前节点的子节点的累加
		step++
	}
	if minPrice == math.MaxInt64 {
		return -1
	}
	return minPrice
}

func QueuePop(queue [][]int) ([]int, [][]int) {
	if len(queue) == 0 {
		return nil, queue
	}
	if len(queue) == 1 {
		return queue[0], queue[:0]
	}
	node := queue[0]
	queue = queue[1:]
	return node, queue
}
