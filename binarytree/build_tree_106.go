package binarytree

func buildTreeByInAndPost(in []int, post []int) *TreeNode {
	if len(in) == 0 || len(post) == 0 {
		return nil
	}
	nodeVal := post[len(post)-1]
	node := &TreeNode{
		Val: nodeVal,
	}
	inNodeIndex := 0
	for i := 0; i < len(in); i++ {
		if in[i] == nodeVal {
			inNodeIndex = i
			break
		}
	}
	inLeft := in[:inNodeIndex]
	node.Left = buildTree(inLeft, post[:len(inLeft)])
	node.Right = buildTree(in[inNodeIndex + 1:], post[len(inLeft) : len(post) - 1])
	return node
}
