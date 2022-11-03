package binarytree

import (
	"algorithm/common"
	"math"
)

// 返回二叉搜索树任意俩节点的差的绝对值的最小值
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	preInGetMinimumDifference = nil
	minInGetMinimumDifference = math.MaxInt64
	getMinimumDifferenceInRecursion(root)
	return minInGetMinimumDifference
}

var preInGetMinimumDifference *TreeNode
var minInGetMinimumDifference int
func getMinimumDifferenceInRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	getMinimumDifferenceInRecursion(root.Left)
	if preInGetMinimumDifference != nil {
		val := common.Abs(root.Val - preInGetMinimumDifference.Val)
		if val < minInGetMinimumDifference {
			minInGetMinimumDifference = val
		}
	}
	preInGetMinimumDifference = root
	getMinimumDifferenceInRecursion(root.Right)
}

