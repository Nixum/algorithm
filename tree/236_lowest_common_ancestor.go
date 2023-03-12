package tree

// 找出p，q的最近公共祖先
// 后序遍历，才能知道左右子树是否包含p或q
// 场景只有两种，p，q在左右两边，那当前root就是最近公共祖先
// 如果p，q在同一边，那最近公共祖先不是p就是q
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// 左子树是否包含p或q
	left := lowestCommonAncestor(root.Left, p, q)
	// 右子树是否包含p或q
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil && right == nil {
		return left
	}
	if left == nil && right != nil {
		return right
	}
	return nil
}
