package binarytree

func inorderTraversal(root *TreeNode) []int {
	inorderRes = make([]int, 0)
	TraversalByInorder(root)
	return inorderRes
}

var inorderRes []int
func TraversalByInorder(root *TreeNode) {
	if root == nil {
		return
	}
	TraversalByInorder(root.Left)
	inorderRes = append(inorderRes, root.Val)
	TraversalByInorder(root.Right)
}

func TraversalByInorderInIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	inorderRes = make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			inorderRes = append(inorderRes, cur.Val)
			cur = cur.Right
		}
	}
	return inorderRes
}

func TraversalByInorderInIteration2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	inorderRes = make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			node = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			inorderRes = append(inorderRes, node.Val)
		}
	}
	return inorderRes
}