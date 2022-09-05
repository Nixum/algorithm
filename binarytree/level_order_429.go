package binarytree

func levelOrder429(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
return res
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	var node *Node
	for len(queue) != 0 {
		size := len(queue)
		levelRes := make([]int, 0)
		for i := 0; i < size; i++ {
			node, queue = QueueNodePop(queue)
			levelRes = append(levelRes, node.Val)
			for j, item := range node.Children {
				if item != nil {
					queue = append(queue, node.Children[j])
				}
			}
		}
		res = append(res, levelRes)
	}
	return res
}
