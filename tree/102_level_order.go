package tree


func levelOrder(root *TreeNode) [][]int {
	levelOrderRes = make([][]int, 0)
	levelOrderInRecursion(root, 0)
	return levelOrderRes
}

var levelOrderRes [][]int
func levelOrderInRecursion(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	depth++
	if len(levelOrderRes) < depth {
		item := make([]int, 0)
		levelOrderRes = append(levelOrderRes, item)
	}
	levelOrderRes[depth - 1] = append(levelOrderRes[depth - 1], root.Val)
	levelOrderInRecursion(root.Left, depth)
	levelOrderInRecursion(root.Right, depth)
}

func levelOrderInIteration(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var node *TreeNode
	for len(queue) != 0 {
		item := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			item = append(item, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, item)
	}
	return res
}
