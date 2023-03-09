package main

import "algorithm/binarytree"

func main() {

}

func lowestCommonAncestor(root, p, q *binarytree.TreeNode) *binarytree.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right:= lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil && right != nil {
		return right
	} else if left != nil && right == nil {
		return left
	}
	return nil
}
