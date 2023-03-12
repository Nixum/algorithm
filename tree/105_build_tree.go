package tree

func buildTreeByInAndPre(pre []int, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	node := &TreeNode{
		Val: pre[0],
	}
	inIndex := 0
	for i := 0; i < len(in); i++ {
		if in[i] == pre[0] {
			inIndex = i
			break
		}
	}
	inLeft := in[:inIndex]
	// 这样写，pre数组容易越界
	//node.Left = buildTreeByInAndPre(pre[1 : len(inLeft)], inLeft)
	//node.Right = buildTreeByInAndPre(pre[len(inLeft) + 1:], in[inIndex + 1 :])
	inRight := in[inIndex + 1 :]
	node.Left = buildTreeByInAndPre(pre[len(pre) - len(inLeft) - len(inRight) : len(pre) - len(inRight)], inLeft)
	node.Right = buildTreeByInAndPre(pre[len(pre) - len(inRight):], inRight)
	return node
}
