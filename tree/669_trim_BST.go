package tree

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		right := trimBST(root.Right, low, high)
		return right
	}
	if root.Val > high {
		left := trimBST(root.Left, low, high)
		return left
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
