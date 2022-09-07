package binarytree

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	res = make([]string, 0)
	backTrackingInBinaryTreePaths(root, "")
	return res
}

var res []string
func backTrackingInBinaryTreePaths(root *TreeNode, path string) {
	if root.Left == nil && root.Right == nil {
		path += fmt.Sprintf("%d", root.Val)
		res = append(res, path)
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