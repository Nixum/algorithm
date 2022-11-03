package binarytree

func preorderTraversal(root *TreeNode) []int {
	preorderRes = make([]int, 0)
	TraversalByPreorder(root)
	return preorderRes
}

var preorderRes []int
func TraversalByPreorder(root *TreeNode) {
	if root == nil {
		return
	}
	preorderRes = append(preorderRes, root.Val)
	TraversalByPreorder(root.Left)
	TraversalByPreorder(root.Right)
}

func TraversalByPreorderInIteration(root *TreeNode) []int {
	preorderRes = make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		preorderRes = append(preorderRes, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return preorderRes
}

// 统一风格迭代
func TraversalByPreorderInIteration2(root *TreeNode) []int {
	preorderRes = make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node, nil)
		} else {
			node = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			preorderRes = append(preorderRes, node.Val)
		}
	}
	return preorderRes
}
