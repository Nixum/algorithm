package tree

// 二叉搜索树中找出众数
func findMode(root *TreeNode) []int {
	countInFindMode = 0
	maxCountInFindMode = 0
	preInFindMode = nil
	resInFindMode = make([]int, 0)
	findModeInRecursion(root)
	return resInFindMode
}

var countInFindMode int
var maxCountInFindMode int
var preInFindMode *TreeNode
var resInFindMode []int
func findModeInRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	findModeInRecursion(root.Left)
	if preInFindMode == nil {
		countInFindMode = 1
	} else if preInFindMode.Val == root.Val {
		countInFindMode++
	} else if preInFindMode.Val != root.Val {
		countInFindMode = 1
	}
	if countInFindMode == maxCountInFindMode {
		resInFindMode = append(resInFindMode, root.Val)
	}
	if countInFindMode > maxCountInFindMode {
		maxCountInFindMode = countInFindMode
		resInFindMode = resInFindMode[:0]
		resInFindMode = append(resInFindMode, root.Val)
	}
	preInFindMode = root
	findModeInRecursion(root.Right)
}