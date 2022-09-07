package binarytree

// 左叶子的和
func sumOfLeftLeaves(root *TreeNode) int {
	sumOfLeftLeaf = 0
	sumOfLeftLeavesInRecursion(root, false)
	return sumOfLeftLeaf
}

var sumOfLeftLeaf int
func sumOfLeftLeavesInRecursion(root *TreeNode, isLeft bool) {
	if isLeft && root.Left == nil && root.Right == nil {
		sumOfLeftLeaf += root.Val
		return
	}
	if root.Left != nil {
		sumOfLeftLeavesInRecursion(root.Left, true)
	}
	if root.Right != nil {
		sumOfLeftLeavesInRecursion(root.Right, false)
	}
	return
}

func sumOfLeftLeaves2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := sumOfLeftLeaves2(root.Left)
	rightSum := sumOfLeftLeaves2(root.Right)
	var midVal int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		midVal = root.Left.Val
	}
	return leftSum + rightSum + midVal
}
