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

func QueuePop(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(queue) == 0 {
		return nil, queue
	}
	if len(queue) == 1 {
		return queue[0], queue[:0]
	}
	node := queue[0]
	queue = queue[1:]
	return node, queue
}