package binarytree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums))
}

func build(nums []int, nodeIndex int, length int) *TreeNode {
	if nodeIndex >= length || nums[nodeIndex] == -1 {
		return nil
	}
	node := &TreeNode{
		Val:   nums[nodeIndex],
		Left:  build(nums, 2 * nodeIndex + 1, length),
		Right: build(nums, 2 * nodeIndex + 2, length),
	}
	return node
}