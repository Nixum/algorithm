package binarytree

// 判断二叉树是不是二叉搜索树
// 左子树的最大节点值小于当前节点，当前节点要小于右边子树最小节点
// 所以中序遍历，且保存前一个节点进行比较即可
// 解法2：递归判断是否是递增
var preNodeInIsValidBST *TreeNode
func isValidBST(root *TreeNode) bool {
	preNodeInIsValidBST = nil
	return isValidBSTInRecursion(root)
}

func isValidBSTInRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftRes := isValidBSTInRecursion(root.Left)
	if preNodeInIsValidBST != nil && root.Val <= preNodeInIsValidBST.Val {
		return false
	}
	preNodeInIsValidBST = root
	rightRes := isValidBSTInRecursion(root.Right)
	return leftRes && rightRes
}


// 验证树是否是二叉搜索树
func isValidBST1(root *TreeNode) bool {
	incrInIsValidBST = make([]int, 0)
	inOrderInIsValidBST(root)
	for i := 1; i < len(incrInIsValidBST); i++ {
		if incrInIsValidBST[i] <= incrInIsValidBST[i - 1] {
			return false
		}
	}
	return true
}

// 解法1：如果是二叉搜索树，那中序遍历得到的序列就是递增的
// 判断是否是递增就能知道是不是二叉搜索树了
var incrInIsValidBST []int
func inOrderInIsValidBST(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderInIsValidBST(root.Left)
	incrInIsValidBST = append(incrInIsValidBST, root.Val)
	inOrderInIsValidBST(root.Right)
}

// 错误的思路，因为这样只判断了当前节点和其左右节点，
// 比如5，4，6，nil，nil，3，7，每个节点的当前判断都是true，但实际上并不符合
func isValidBSTInRecursion1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := isValidBSTInRecursion1(root.Left)
	right := isValidBSTInRecursion1(root.Right)
	isValid := false
	if root.Left != nil && root.Right == nil && root.Left.Val < root.Val {
		isValid = true
	}
	if root.Right != nil && root.Left == nil && root.Right.Val > root.Val {
		isValid = true
	}
	if root.Left != nil && root.Right != nil &&
		root.Left.Val < root.Val && root.Val < root.Right.Val {
		isValid = true
	}
	if root.Left == nil && root.Right == nil {
		isValid = true
	}
	return left && right && isValid
}