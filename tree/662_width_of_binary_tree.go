package tree

import "algorithm/common"

// dfs + 节点编号
var minValue map[int]int // 用来保存当前层的最小值，dfs的时候，遍历的第一个节点一定是最小值
var resInWidthOfBinaryTree int
func widthOfBinaryTree(root *TreeNode) int {
	resInWidthOfBinaryTree = 0
	minValue =  make(map[int]int)
	dfsInWidthOfBinaryTree(root, 1, 0)
	return resInWidthOfBinaryTree
}

func dfsInWidthOfBinaryTree(root *TreeNode, nodeIndex int, level int) {
	if root == nil {
		return
	}
	if _, exist := minValue[level]; !exist {
		minValue[level] = nodeIndex
	}
	resInWidthOfBinaryTree = common.Max(resInWidthOfBinaryTree, nodeIndex - minValue[level]+1)
	dfsInWidthOfBinaryTree(root, 2 * nodeIndex, level+1)
	dfsInWidthOfBinaryTree(root, 2 * nodeIndex + 1, level+1)
}

// bfs + 节点编号
func widthOfBinaryTree2(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, &TreeNode{
		Val:   1,
		Left:  root.Left,
		Right: root.Right,
	})
	var node *TreeNode
	res := 0
	for len(queue) > 0 {
		size := len(queue)
		startIndex := -1
		endIndex := -1
		for i := 0; i < size; i++ {
			node, queue = QueuePop(queue)
			// 记录每层的最后一个编号
			endIndex = node.Val
			// 记录每层的第一个编号
			if startIndex == -1 {
				startIndex = node.Val
			}
			if node.Left != nil {
				queue = append(queue, &TreeNode{
					Val:   node.Val * 2,
					Left:  node.Left.Left,
					Right: node.Left.Right,
				})
			}
			if node.Right != nil {
				queue = append(queue, &TreeNode{
					Val:   node.Val * 2 + 1,
					Left:  node.Right.Left,
					Right: node.Right.Right,
				})
			}
		}
		res = common.Max(res, endIndex - startIndex + 1)
	}
	return res
}
