package binarytree

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBySortedArrayToBST(nums, 0, len(nums) - 1)
}

func buildBySortedArrayToBST(nums []int, l int, h int) *TreeNode {
	if l > h {
		return nil
	}
	mid := l + (h - l) / 2
	node := &TreeNode{
		Val:   nums[mid],
		Left:  buildBySortedArrayToBST(nums, l, mid - 1),
		Right: buildBySortedArrayToBST(nums, mid + 1, h),
	}
	return node
}
