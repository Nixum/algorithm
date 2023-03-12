package tree

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	resInBinaryTreePaths = make([]string, 0)
	backTrackingInBinaryTreePaths(root, "")
	return resInBinaryTreePaths
}

var resInBinaryTreePaths []string
func backTrackingInBinaryTreePaths(root *TreeNode, path string) {
	if root.Left == nil && root.Right == nil {
		path += fmt.Sprintf("%d", root.Val)
		resInBinaryTreePaths = append(resInBinaryTreePaths, path)
		return
	}
	path += fmt.Sprintf("%d->", root.Val)
	if root.Left != nil {
		backTrackingInBinaryTreePaths(root.Left, path)
	}
	if root.Right != nil {
		backTrackingInBinaryTreePaths(root.Right, path)
	}
}