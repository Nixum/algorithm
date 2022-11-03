package binarytree

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	head := root
	queue := make([]*Node, 0)
	queue = append(queue, root)
	var p *Node
	for len(queue) != 0 {
		size := len(queue)
		h := new(Node)
		for i := 0; i < size; i++ {
			p, queue = QueueNodePop(queue)
			h.Next = p
			h = p
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
	}
	return head
}
