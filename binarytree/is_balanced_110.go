package binarytree

import (
	"algorithm/common"
	"math"
)

func isBalanced(root *TreeNode) bool {
	if isBalancedGetHeight(root) == -1 {
		return false
	}
	return true
}

func isBalancedGetHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lH := isBalancedGetHeight(root.Left)
	if lH == -1 {
		return -1
	}
	rH := isBalancedGetHeight(root.Right)
	if rH == -1 {
		return -1
	}
	if math.Abs(float64(lH - rH)) > 1 {
		return -1
	}
	return 1 + common.Max(lH, rH)
}
