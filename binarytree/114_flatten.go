package binarytree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flattenInRecursion(root)
}

func flattenInRecursion(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := flattenInRecursion(root.Left)
	right := flattenInRecursion(root.Right)
	tmp := left
	root.Right = tmp
	root.Left = nil
	if tmp == nil {
		root.Right = right
	} else {
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = right
	}
	return root
}
