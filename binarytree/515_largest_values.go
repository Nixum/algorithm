package binarytree

import "math"

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var node *TreeNode
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		max := math.MinInt64
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
