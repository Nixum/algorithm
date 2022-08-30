package binarytree

func inorderTraversal(root *TreeNode) []int {
	inorderRes = make([]int, 0)
	TraversalByInorder(root)
	return inorderRes
}

var inorderRes []int
func TraversalByInorder(root *TreeNode) {
	if root == nil {
		return
	}
	TraversalByInorder(root.Left)
	inorderRes = append(inorderRes, root.Val)
	TraversalByInorder(root.Right)
}
