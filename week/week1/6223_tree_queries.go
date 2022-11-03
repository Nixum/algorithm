package week1

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 思路：
// 1. 先跑一遍DFS算出每棵子树的高度
// 2. DFS遍历每个节点，算出如果移除当前节点后树的最大高度
// 计算当前节点被移除后的最大高度：
// 左边：max(当前深度+右子树高度, 删除当前子树后的剩余高度)
// 右边：max(当前深度+左子树高度, 删除当前子树后的剩余高度)
func treeQueries(root *TreeNode, queries []int) []int {
	heightMap := make(map[*TreeNode]int)
	traverse(root, heightMap)
	resMap := make(map[int]int)
	traverseCal(root, -1, 0, resMap, heightMap)
	res := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		res = append(res, resMap[queries[i]])
	}
	return res
}

func traverse(root *TreeNode, hMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	hMap[root] = 1 + max(traverse(root.Left, hMap), traverse(root.Right, hMap))
	return hMap[root]
}

// resH的含义：删除当前子树后剩余部分树的高度
func traverseCal(root *TreeNode, depth int, resH int, restHeight map[int]int, hMap map[*TreeNode]int) {
	if root == nil {
		return
	}
	depth++
	restHeight[root.Val] = resH
	// 删除左子树时，最高高度：max(当前高度为resH, 当前高度 + 右子树高度)
	traverseCal(root.Left, depth, max(resH, depth + hMap[root.Right]), restHeight, hMap)
	// 删除右子树时，最高高度：max(当前高度为resH, 当前高度 + 左子树高度)
	traverseCal(root.Right, depth, max(resH, depth + hMap[root.Left]), restHeight, hMap)
}

// 难点：每次遍历完要恢复成初始状态, 但下面这种解法超时了
func treeQueries2(root *TreeNode, queries []int) []int {
	res = make([]int, 0)
	for i := 0; i < len(queries); i++ {
		height := traverse2(copyTree(root), queries[i])
		res = append(res, height - 1)
	}
	return res
}

var res []int

func traverse2(root *TreeNode, n int) int {
	if root == nil || root.Val == n {
		return 0
	}
	return 1 + max(traverse2(root.Left, n), traverse2(root.Right, n))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head := &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
	return head
}