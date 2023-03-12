package tree

func pathSumIII(root *TreeNode, targetSum int) int {
	sum2NumMap := make(map[int]int)
	sum2NumMap[0] = 1
	return prefixSumInPathSumInIII(root, 0, targetSum, sum2NumMap)
}

func prefixSumInPathSumInIII(root *TreeNode, sum, targetSum int, sum2NumMap map[int]int) int {
	if root == nil {
		return 0
	}
	res := 0
	// 计算前缀和
	sum += root.Val
	// 两节点的前缀和 = 两节点的前缀和之差
	// prefix(A) - prefix(B) = targetSum
	// map的value存对应前缀和只差的数量
	res += sum2NumMap[sum-targetSum]
	// 计算前缀和的数量
	sum2NumMap[sum] = sum2NumMap[sum]+1

	// 计算 子节点的前缀和之差 的路径数量
	leftSum := prefixSumInPathSumInIII(root.Left, sum, targetSum, sum2NumMap)
	rightSum := prefixSumInPathSumInIII(root.Right, sum, targetSum, sum2NumMap)

	res += leftSum + rightSum
	// 恢复回原状态
	sum2NumMap[sum] = sum2NumMap[sum]-1
	return res
}

// dfs，从上到下判断以当前节点为根节点的路径和是否等于target
func pathSumIII2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := backtrackingInPathSumInIII(root, targetSum)
	res += pathSumIII2(root.Left, targetSum)
	res += pathSumIII2(root.Right, targetSum)
	return res
}

// 计算当前节点的路径和
func backtrackingInPathSumInIII(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val == targetSum {
		res++
	}
	res += backtrackingInPathSumInIII(root.Left, targetSum - root.Val)
	res += backtrackingInPathSumInIII(root.Right, targetSum - root.Val)
	return res
}

// ------------------ 下面的思路有起点问题，应该每个节点重新计算路径和
func pathSumIII3(root *TreeNode, targetSum int) int {
	resInIII = 0
	backtrackingInPathSumInIII3(root, targetSum, targetSum)
	return resInIII
}

var resInIII int
func backtrackingInPathSumInIII3(root *TreeNode, targetSum int, originalSum int) {
	if root == nil {
		return
	}
	val := targetSum - root.Val
	if targetSum - root.Val < 0 {
		val = originalSum
	}
	if 0 == val {
		resInIII++
		val = originalSum - root.Val
	}
	backtrackingInPathSumInIII3(root.Left, val, originalSum)
	backtrackingInPathSumInIII3(root.Right, val, originalSum)
}
