package dp

import (
	"algorithm/common"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][][]int)
	for i := 0; i < len(flights); i++ {
		// 记录谁指向该节点
		graph[flights[i][1]] = append(graph[flights[i][1]], []int{flights[i][0], flights[i][2]})
	}
	// k转成边
	k++
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k + 1)
		for j := 0; j < k + 1; j++ {
			dp[i][j] = -2
		}
	}
	return dpInFindCheapestPrice(src, dst, k, dp, graph)
}

// dp[s][k]：从src出发，k步之内到达s的最小权重
// 递推公式：dp[dst][k] = min(dp[s1][k-1] + w1, dp[s2][k-1] + w2)
func dpInFindCheapestPrice(src int, s int, k int, dp [][]int, graph map[int][][]int) int {
	if s == src {
		return 0
	}
	if k == 0 {
		return -1
	}
	if dp[s][k] != -2 {
		return dp[s][k]
	}
	res := math.MaxInt64
	if _, exist := graph[s]; exist {
		for _, next := range graph[s] {
			// 从后往前推
			nextPrice := dpInFindCheapestPrice(src, next[0], k - 1, dp, graph)
			if nextPrice != -1 {
				res = common.Min(res, nextPrice + next[1])
			}
		}
	}
	if res == math.MaxInt64 {
		dp[s][k] = -1
	} else {
		dp[s][k] = res
	}
	return dp[s][k]
}
