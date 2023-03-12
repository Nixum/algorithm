package tree

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

// 其他写法，也能ac
func flatten1(root *TreeNode)  {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tmp := root.Right
	root.Right = root.Left
	leftTmp := root.Left
	// left要置为nil！！
	root.Left = nil
	for leftTmp != nil && leftTmp.Right != nil {
		leftTmp = leftTmp.Right
	}
	if leftTmp != nil {
		leftTmp.Right = tmp
	} else {
		root.Right = tmp
	}
}