package binarytree

func numTrees(n int) int {
	memo = make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n + 1)
	}
	return countInNumTrees(1, n)
}

var memo [][]int
func countInNumTrees(left, right int) int {
	// 空节点也算一种结果，所以返回 1
	if left > right {
		return 1
	}
	if memo[left][right] != 0 {
		return memo[left][right]
	}
	res := 0
	for i := left; i <= right; i++ {
		numL := countInNumTrees(left, i - 1)
		numR := countInNumTrees(i + 1, right)
		res += numL * numR
	}
	memo[left][right] = res
	return res
}
