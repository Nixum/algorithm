package binarytree

import "math"

func constructMaximumBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	maxVal := math.MinInt64
	maxValIndex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxValIndex = i
		}
	}
	node := &TreeNode{
		Val: maxVal,
	}
	node.Left = constructMaximumBinaryTree(arr[:maxValIndex])
	node.Right = constructMaximumBinaryTree(arr[maxValIndex + 1:])
	return node
}