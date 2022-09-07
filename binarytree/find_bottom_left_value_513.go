package binarytree

import "math"


var bottomLeftValue int
var heightInFindBottomLeftValue int
func findBottomLeftValue(root *TreeNode) int {
	bottomLeftValue = 0
	heightInFindBottomLeftValue = 0
	findBottomLeftValueInRecursion(root, 0)
	return bottomLeftValue
}

func findBottomLeftValueInRecursion(root *TreeNode, leftHeight int) {
	if root.Left == nil && root.Right == nil {

		return
	}
	return
}

func findBottomLeftValue2(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	leftV := math.MinInt64
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			if i == 0 {
				leftV = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return leftV
}
