package binarytree

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return make([]*TreeNode, 0)
	}
	return buildInGenerateTrees(1, n)
}

func buildInGenerateTrees(left, right int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if left > right {
		res = append(res, nil)
		return res
	}
	// 穷举根节点所有可能
	for i := left; i <= right; i++ {
		leftNodes := buildInGenerateTrees(left, i - 1)
		rightNodes := buildInGenerateTrees(i + 1, right)
		for j := 0; j < len(leftNodes); j++ {
			for k := 0; k < len(rightNodes); k++ {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  leftNodes[j],
					Right: rightNodes[k],
				})
			}
		}
	}
	return res
}
