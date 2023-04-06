package monotonic_stack

import (
	"algorithm/common"
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
	max := 0
	row := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				row[j] += 1
			} else {
				row[j] = 0
			}
		}
		max = common.Max(max, calcInMaximalRectangle(row))
	}
	return max
}

// 单调栈，栈内单调递增，当栈顶元素比当前大，则出栈，直到栈顶元素比当前小
func calcInMaximalRectangle(row []int) int {
	stack := make([]int, 0)
	row = append(row, 0)
	max := row[0]
	stack = append(stack, 0)
	for i := 1; i < len(row); i++ {
		for len(stack) > 0 && row[stack[len(stack)-1]] > row[i] {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prePre := -1
			if len(stack) > 0 {
				prePre = stack[len(stack)-1]
			}
			fmt.Println(prePre, pre, i, row[pre], (i - prePre - 1) * row[pre])
			// pre是栈顶元素，i是右边第一个比他小的，prePre是左边第一个比他小的
			// 计算的时候，是不包含prePre和i所在的高度的，仅用他们的下标来计算
			max = common.Max(max, (i - prePre - 1) * row[pre])
		}
		fmt.Println()
		stack = append(stack, i)
	}
	return max
}