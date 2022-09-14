package binarytree

// BST中删除节点，要分情况，当前节点是否为空，左右节点是否为空，
// 当左右节点均不为空时，需要把左节点挂在右子树最左节点的左节点上
// 如果不是BST的二叉树删除节点，只需要把要删除的节点和叶子节点交换删除即可
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right != nil && root.Left == nil {
			return root.Right
		} else if root.Left == nil && root.Right != nil {
			return root.Left
		} else {
			tmp := root.Right
			for tmp.Left != nil {
				tmp = tmp.Left
			}
			tmp.Left = root.Left
			return root.Right
		}
	}
	if key > root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key < root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
