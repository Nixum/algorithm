package binarytree

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
