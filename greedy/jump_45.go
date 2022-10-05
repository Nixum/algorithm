package greedy

import (
	"algorithm/common"
)

func jump(nums []int) int {
	res := 0
	curPath := 0
	nextPath := 0
	for i := 0; i < len(nums); i++ {
		nextPath = common.Max(i + nums[i], nextPath)
		// 到达当前最远覆盖距离的下标
		if curPath == i {
			// 不是终点
			if curPath != len(nums) - 1 {
				res++
				curPath = nextPath
				// 如果下一步能到终点就不需要继续走了
				if nextPath >= len(nums) - 1 {
					break
				}
			} else {
				break
			}
		}
	}
	return res
}

func jump2(nums []int) int {
	res := 0
	curPath := 0
	nextPath := 0
	// 关键点， len(nums) - 1，因为题目是一定可以跳到终点，所以到倒数第二个后再走一步即可
	for i := 0; i < len(nums) - 1; i++ {
		nextPath = common.Max(i + nums[i], curPath)
		// 一直往最大距离移动
		if i == curPath {
			curPath = nextPath
			res++
		}
	}
	return res
}
