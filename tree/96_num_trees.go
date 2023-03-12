package tree

// 可以和95题一起看
func numTrees(n int) int {
	memo = make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n + 1)
	}
	return countInNumTrees(1, n)
}

var memo [][]int
// 闭区间[left, fight]的数组能组成count(left, right)种BST
func countInNumTrees(left, right int) int {
	// 空节点也算一种结果，所以返回 1
	if left > right {
		return 1
	}
	if memo[left][right] != 0 {
		return memo[left][right]
	}
	res := 0
	// 每个节点都要作为根节点
	for i := left; i <= right; i++ {
		// 分为左子树
		numL := countInNumTrees(left, i - 1)
		// 分为右子树
		numR := countInNumTrees(i + 1, right)
		// 左右子树的组合数乘积是BST的总数
		res += numL * numR
	}
	memo[left][right] = res
	return res
}
