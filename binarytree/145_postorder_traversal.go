package binarytree

func postorderTraversal(root *TreeNode) []int {
	postorderRes = make([]int, 0)
	TraversalByPostorder(root)
	return postorderRes
}

var postorderRes []int
func TraversalByPostorder(root *TreeNode) {
	if root == nil {
		return
	}
	TraversalByPostorder(root.Left)
	TraversalByPostorder(root.Right)
	postorderRes = append(postorderRes, root.Val)
}

// 先序遍历迭代法的反转
func TraversalByPostorderInIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	postorderRes = make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		postorderRes = append(postorderRes, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	i := 0
	j := len(postorderRes) - 1
	for  i <= j {
		postorderRes[i], postorderRes[j] = postorderRes[j], postorderRes[i]
		i++
		j--
	}
	return postorderRes
}


func TraversalByPostorderInIteration2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	postorderRes = make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if node != nil {
			stack = append(stack, node, nil)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			node = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			postorderRes = append(postorderRes, node.Val)
		}
	}
	return postorderRes
}