package tree

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			// 每层的最后一个元素放入
			if i == size - 1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
