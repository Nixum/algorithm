package binarytree

func convertBST(root *TreeNode) *TreeNode {
	sumInConvertBST = 0
	convertBSTInRecursion(root)
	return root
}

var sumInConvertBST int
func convertBSTInRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTInRecursion(root.Right)
	sumInConvertBST += root.Val
	root.Val = sumInConvertBST
	convertBSTInRecursion(root.Left)
}
