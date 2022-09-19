package binarytree

func bstToGst(root *TreeNode) *TreeNode {
	sumInBstToGst = 0
	bstToGstInRecursion(root)
	return root
}

var sumInBstToGst int
func bstToGstInRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	bstToGstInRecursion(root.Right)
	sumInBstToGst += root.Val
	root.Val = sumInBstToGst
	bstToGstInRecursion(root.Left)
	return
}
