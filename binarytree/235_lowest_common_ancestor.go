package binarytree

// 二叉搜索树中的最近公共祖先，利用有序性，只需要遍历一边
func lowestCommonAncestorInBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestorInBST(root.Left, p, q)
		if left != nil {
			return left
		}
	} else if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestorInBST(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	// 当 root.Val 的值在 [p.Val, q.Val] 中间(左右闭区间)时，就是最近公共祖先
	return root
}
