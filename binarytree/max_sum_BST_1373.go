package binarytree

import (
	"algorithm/common"
	"math"
)

// 1. 判断左右子树是否是BST
// 2. 左子树的最大值和右子树的最小值
// 3. 左右子树的节点值之和
func maxSumBST(root *TreeNode) int {
	maxSumInMaxSumBST = 0
	maxSumBSTInRecursion(root)
	return maxSumInMaxSumBST
}


// res结果数组，0：root为根的二叉树是否是BST，1是0不是
// 1：root为根的二叉树所有节点最小值，2：最大值
// 3：root为根的二叉树的所有节点之和
var maxSumInMaxSumBST int
func maxSumBSTInRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{1, math.MaxInt64, math.MinInt64, 0}
	}
	leftRes := maxSumBSTInRecursion(root.Left)
	rightRes := maxSumBSTInRecursion(root.Right)
	
	if leftRes[0] == 1 && rightRes[0] == 1 &&
		root.Val > leftRes[2] && root.Val < rightRes[1] {
		minVal := common.Min(leftRes[1], root.Val)
		maxVal := common.Max(rightRes[2], root.Val)
		sum :=  leftRes[3] + rightRes[3] + root.Val
		if sum > maxSumInMaxSumBST {
			maxSumInMaxSumBST = sum
		}
		return []int{1, minVal, maxVal, sum}
	} else {
		return []int{0, 0, 0, 0}
	}
}
