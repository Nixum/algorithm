package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	c := len(matrix[0])-1
	r := 0
	for c >= 0 && r < len(matrix) {
		if matrix[r][c] > target {
			c--
		} else if matrix[r][c] < target {
			r++
		} else {
			return true
		}
	}
	return false
}

// 思路是对的，
// 但是这种写法不行，因为取的begin1=mid，可能会导致死循环
// 比如[1,1] target=1的场景
// 如果取 begin1=mid+1，又有可能导致数组越界
func searchMatrix2(matrix [][]int, target int) bool {
	begin1, end1 := 0, len(matrix[0])-1
	for begin1 < end1 {
		mid := begin1 + (end1 - begin1) / 2
		if matrix[mid][0] < target {
			begin1 = mid
		} else if matrix[mid][0] > target {
			end1 = mid - 1
		} else {
			return true
		}
	}
	begin2, end2 := 0, len(matrix[begin1])-1
	for begin2 < end2 {
		mid := begin2 + (end2 - begin2) / 2
		if matrix[begin1][mid] < target {
			begin2 = mid
		} else if matrix[begin1][mid] > target {
			end2 = mid - 1
		} else {
			return true
		}
	}
	return matrix[begin1][begin2] == target
}
