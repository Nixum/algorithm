package tree

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, float64(sum) / float64(size))
	}
	return res
}
