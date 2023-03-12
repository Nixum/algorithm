package tree

func kthSmallest(root *TreeNode, k int) int {
	kInKthSmallest = 0
	resInKthSmallest = 0
	kthSmallestInRecursion(root, k)
	return resInKthSmallest
}

var kInKthSmallest int
var resInKthSmallest int
func kthSmallestInRecursion(root *TreeNode, k int) {
	if root == nil {
		return
	}
	kthSmallestInRecursion(root.Left, k)
	kInKthSmallest++
	if kInKthSmallest == k {
		resInKthSmallest = root.Val
		return
	}
	kthSmallestInRecursion(root.Right, k)
}
