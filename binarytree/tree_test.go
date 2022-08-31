package binarytree

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(preorderTraversal(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPreorderInIteration(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPreorderInIteration2(root))
}

func TestPostorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(postorderTraversal(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPostorderInIteration(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPostorderInIteration2(root))
}

func TestInorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(inorderTraversal(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByInorderInIteration(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByInorderInIteration2(root))
}

func TestLevelOrder(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(levelOrder(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(levelOrder(root))

	root = CreateBinaryTree([]int{})
	fmt.Println(levelOrder(root))
}

func TestLevelOrderBottom(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(levelOrderBottom(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(levelOrderBottom(root))

	root = CreateBinaryTree([]int{})
	fmt.Println(levelOrderBottom(root))
}
