package array

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if matrix == nil || len(matrix) == 0 {
		return res
	}
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	num := len(matrix) * len(matrix[0])
	for num > 0 {
		for i := left; i <= right && num > 0; i++ {
			res = append(res, matrix[top][i])
			num--
		}
		top++
		for i := top; i <= bottom && num > 0; i++ {
			res = append(res, matrix[i][right])
			num--
		}
		right--
		for i := right; i >= left && num > 0; i-- {
			res = append(res, matrix[bottom][i])
			num--
		}
		bottom--
		for i := bottom; i >= top && num > 0; i-- {
			res = append(res, matrix[i][left])
			num--
		}
		left++
	}
	return res
}
