package binarytree

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(preorderTraversal(root))
}

func TestPostorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(postorderTraversal(root))
}

func TestInorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(inorderTraversal(root))
}
