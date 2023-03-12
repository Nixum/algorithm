package tree

func hasPathSum(root *TreeNode, t int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == t {
		return true
	}
	leftRes := hasPathSum(root.Left, t - root.Val)
	rightRes := hasPathSum(root.Right, t - root.Val)
	return leftRes || rightRes
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumInRecursion(root, 0, targetSum)
}

func hasPathSumInRecursion(root *TreeNode, sum int, targetSum int) bool {
	if root == nil {
		return false
	}
	// 关键在这，不能在root==nil的时候比较，否则无法处理 {1, 2} t=1的场景
	if root.Left == nil && root.Right == nil &&
		sum + root.Val == targetSum {
		return true
	}
	leftRes := hasPathSumInRecursion(root.Left, sum + root.Val, targetSum)
	if leftRes {
		return true
	}
	rightRes := hasPathSumInRecursion(root.Right, sum + root.Val, targetSum)
	if rightRes {
		return true
	}
	return false
}
