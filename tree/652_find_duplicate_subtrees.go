package tree

import "fmt"

// 返回重复的子树的根节点
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	path2FreqInFindDuplicateSubtrees = make(map[string]int)
	resInFindDuplicateSubtrees = make([]*TreeNode, 0)
	findDuplicateSubtreesInRecursion(root)
	return resInFindDuplicateSubtrees
}

var path2FreqInFindDuplicateSubtrees map[string]int
var resInFindDuplicateSubtrees []*TreeNode
func findDuplicateSubtreesInRecursion(root *TreeNode) string {
	if root == nil {
		return ""
	}
	left := findDuplicateSubtreesInRecursion(root.Left)
	right := findDuplicateSubtreesInRecursion(root.Right)
	path := fmt.Sprintf("%d,%s,%s", root.Val, left, right)
	// 只能等于1，不然会重复
	if path2FreqInFindDuplicateSubtrees[path] == 1 {
		resInFindDuplicateSubtrees = append(resInFindDuplicateSubtrees, root)
	}
	path2FreqInFindDuplicateSubtrees[path]++
	return path
}
