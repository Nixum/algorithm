package tree

import (
	"algorithm/common"
	"math"
)

func maxPathSum(root *TreeNode) int {
	maxInMaxPathSum = math.MinInt64
	dfsInMaxPathSum(root)
	return maxInMaxPathSum
}

// 思路:
// 定义dfs方法为：返回当前节点的最大路径和
// 当前节点的最大路径和 = 左节点最大路径和 + 右节点最大路径和 + 当前节点的值
// 返回时，由于路径只能选一个方向，所以是选左右节点最大的路径和 + 当前节点
// 由于相加可能存在负数，所以要跟0比较，如果是负数则直接舍弃即可
var maxInMaxPathSum int
func dfsInMaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := dfsInMaxPathSum(root.Left)
	rightMax := dfsInMaxPathSum(root.Right)
	maxInMaxPathSum = common.Max(maxInMaxPathSum, root.Val + leftMax + rightMax)
	nodeMax := common.Max(leftMax, rightMax) + root.Val
	return common.Max(nodeMax, 0)
}