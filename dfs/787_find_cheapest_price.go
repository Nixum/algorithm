package dfs

import "math"

// DFS, 改进版本, 但仍然超时, 时间复杂度O(n^k)
func findCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][][]int)
	for i := 0; i < len(flights); i++ {
		s := flights[i][0]
		d := flights[i][1]
		p := flights[i][2]
		graph[s] = append(graph[s], []int{d, p})
	}
	minPriceInFindCheapestPrice = math.MaxInt64
	handleNextInFindCheapestPrice2(graph, src, dst, k+1, 0)
	if minPriceInFindCheapestPrice == math.MaxInt64 {
		return -1
	}
	return minPriceInFindCheapestPrice
}

func handleNextInFindCheapestPrice2(graph map[int][][]int, src, dst int, k int, price int) {
	if k < 0 {
		return
	}
	if src == dst {
		minPriceInFindCheapestPrice = price
		return
	}
	if _, exist := graph[src]; !exist {
		return
	}
	for _, next := range graph[src] {
		if price + next[1] > minPriceInFindCheapestPrice {
			continue
		}
		handleNextInFindCheapestPrice2(graph, next[0], dst, k - 1, price + next[1])
	}
	return
}

var minPriceInFindCheapestPrice int
// ----------------------------------------
// DFS思路是对的，但是超时了
func findCheapestPrice3(n int, flights [][]int, src int, dst int, k int) int {
	// key是起点，value是[]int的数组，[]int，下标代表的含义：0：终点，1：权重，2：是否被使用，0表示未被使用，1表示被使用
	graph := make(map[int][][]int)
	for i := 0; i < len(flights); i++ {
		s := flights[i][0]
		d := flights[i][1]
		p := flights[i][2]
		graph[s] = append(graph[s], []int{d, p, 0})
	}
	minPriceInFindCheapestPrice = math.MaxInt64
	for _, next := range graph[src] {
		res := handleNextInFindCheapestPrice3(next, graph, dst, k, 0)
		if res < minPriceInFindCheapestPrice {
			minPriceInFindCheapestPrice = res
		}
	}
	if minPriceInFindCheapestPrice == math.MaxInt64 {
		return -1
	}
	return minPriceInFindCheapestPrice
}

func handleNextInFindCheapestPrice3(cur []int, graph map[int][][]int, dst int, k int, price int) int {
	k--
	price = price + cur[1]
	if k < -1 || cur[2] == 1 {
		return math.MaxInt64
	}
	if cur[0] == dst {
		return price
	}
	if _, exist := graph[cur[0]]; !exist {
		return math.MaxInt64
	}
	cur[2] = 1
	for _, next := range graph[cur[0]] {
		if price + next[1] > minPriceInFindCheapestPrice {
			continue
		}
		res := handleNextInFindCheapestPrice3(next, graph, dst, k, price)
		if res < minPriceInFindCheapestPrice {
			minPriceInFindCheapestPrice = res
		}
	}
	cur[2] = 0
	return minPriceInFindCheapestPrice
}
