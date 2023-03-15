package binarysearch

// 时间复杂度O(nlogn)，空间复杂度O(1)
func searchMatrixII(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		begin := 0
		end := len(matrix[i])-1
		for begin <= end {
			mid := begin + (end - begin)/2
			if matrix[i][mid] < target {
				begin = mid + 1
			} else if matrix[i][mid] > target {
				end = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}

// 抽象乘以右上角为根的右子树，如果目标值小于当前值，搜索左子树（向左），否则搜索右子树（向下）
// 时间复杂度O(n)，空间复杂度O(1)
func searchMatrixII2(matrix [][]int, target int) bool {
	r := 0
	c := len(matrix[0])-1
	for r < len(matrix) && c >= 0 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}