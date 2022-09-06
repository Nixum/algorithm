package binarytree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricCompare(root.Left, root.Right)
}

func isSymmetricCompare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
	}
	outSide := isSymmetricCompare(left.Left, right.Right)
	inSide := isSymmetricCompare(left.Right, right.Left)
	return outSide && inSide
}
