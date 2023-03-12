package tree

import "algorithm/common"

// 左右孩子为空时，才是算是低点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			// 注意得在里面判断
			// 如果在外面判断会导致只判断左节点
			// 如果有右节点的左右子节点是空，就会漏判，因为此时已经pop出来了
			if node.Left == nil && node.Right == nil {
				return depth
			}
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

func minDepthInRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepthInRecursion(root.Left)
	rightDepth := minDepthInRecursion(root.Right)
	// 因为是要取最小高度，所以只要有一边是null就能返回了
	if root.Left == nil && root.Right != nil {
		return 1 + rightDepth
	}
	if root.Left != nil && root.Right == nil {
		return 1 + leftDepth
	}
	return 1 + common.Min(leftDepth, rightDepth)
}