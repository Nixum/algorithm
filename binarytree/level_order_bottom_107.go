package binarytree

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		items := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			items = append(items, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, items)
	}
	i := 0
	j := len(res) - 1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
