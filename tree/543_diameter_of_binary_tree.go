package tree

import "algorithm/common"

// 求二叉树的最大直径
// 当前节点的最大直径 = 左节点的最大深度 + 右节点的最大深度
// 当前节点的最大深度 = max(左节点最大深度，右节点最大深度) + 1
// 二叉树的最大直径等于某个节点的最大直径，该节点不一定是根节点
func diameterOfBinaryTree(root *TreeNode) int {
	maxWide = 0
	heightInDiameterOfBinaryTree(root)
	return maxWide
}

var maxWide int = 0
func heightInDiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := heightInDiameterOfBinaryTree(root.Left)
	rightDepth := heightInDiameterOfBinaryTree(root.Right)
	maxWide = common.Max(maxWide, leftDepth + rightDepth)
	return common.Max(leftDepth, rightDepth) + 1
}
