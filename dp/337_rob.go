package dp

import "algorithm/common"

// 树形dp
func rob(root *TreeNode) int {
	res := robInTree(root)
	return common.Max(res[0], res[1])
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 后序遍历，判断要不要抢当前节点
// 如果抢了当前节点，两孩子节点就不能抢了
// 如果不抢当前节点，就要考虑左右节点抢还是不抢，即 取左右节点最大值相加
// 返回值 [2]int, 表示 偷和不偷
func robInTree(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{}
	}
	resLeft := robInTree(root.Left)
	resRight := robInTree(root.Right)
	res := [2]int{
		// 偷
		root.Val + resLeft[1] + resRight[1],
		// 不偷，从左右孩子里取最大值
		common.Max(resLeft[0], resLeft[1]) + common.Max(resRight[0], resRight[1]),
	}
	return res
}