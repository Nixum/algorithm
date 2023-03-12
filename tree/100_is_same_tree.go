package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
