package binarytree

func postorderTraversal(root *TreeNode) []int {
	postorderRes = make([]int, 0)
	TraversalByPostorder(root)
	return postorderRes
}

var postorderRes []int
func TraversalByPostorder(root *TreeNode) {
	if root == nil {
		return
	}
	TraversalByPostorder(root.Left)
	TraversalByPostorder(root.Right)
	postorderRes = append(postorderRes, root.Val)
}
