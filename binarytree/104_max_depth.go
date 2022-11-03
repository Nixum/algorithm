package binarytree

import "algorithm/common"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		size := len(queue)
		depth++
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}


func maxDepthInRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + common.Max(maxDepthInRecursion(root.Left),
		maxDepthInRecursion(root.Right))
}