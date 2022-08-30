package binarytree

func preorderTraversal(root *TreeNode) []int {
	preorderRes = make([]int, 0)
	TraversalByPreorder(root)
	return preorderRes
}

var preorderRes []int
func TraversalByPreorder(root *TreeNode) {
	if root == nil {
		return
	}
	preorderRes = append(preorderRes, root.Val)
	TraversalByPreorder(root.Left)
	TraversalByPreorder(root.Right)
}
