package array

// 参考：https://leetcode.cn/problems/rotate-image/solution/lu-qing-ge-chong-by-pennx-ce3x/
// 通过观察可得，m[i][j] 顺时针选择90度变为：m[j][n-i-1]
// 相当于，先上下对称，m[i][j] 变为 m[n-i-1][j],
// 再主对角线对称，m[i][j] 变为 m[j][i]
// 其他对称，比如左右对称，m[i][j] -> m[i][n-j-1]
// 其他对称，比如副对角线对称，m[i][j] -> m[n-j-1][n-i-1]
func rotate(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m / 2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[m-i-1][j] = matrix[m-i-1][j], matrix[i][j]
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
