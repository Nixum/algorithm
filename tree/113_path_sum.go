package tree

var resInPathSum [][]int
var pathInPathSum []int
func pathSum(root *TreeNode, targetSum int) [][]int {
	resInPathSum = make([][]int, 0)
	pathInPathSum = make([]int, 0)
	if root == nil {
		return resInPathSum
	}
	pathSumInRecursion(root, targetSum)
	return resInPathSum
}

func pathSumInRecursion(root *TreeNode, targetSum int) {
	if root.Left == nil && root.Right == nil &&
		root.Val == targetSum {
		res := make([]int, 0)
		res = append(res, pathInPathSum...)
		res = append(res, root.Val)
		resInPathSum = append(resInPathSum, res)
		return
	}
	pathInPathSum = append(pathInPathSum, root.Val)
	if root.Left != nil {
		pathSumInRecursion(root.Left, targetSum - root.Val)
	}
	if root.Right != nil {
		pathSumInRecursion(root.Right, targetSum - root.Val)
	}
	pathInPathSum = pathInPathSum[:len(pathInPathSum) - 1]
}

