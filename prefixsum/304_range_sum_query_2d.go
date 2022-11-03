package prefixsum

type NumMatrix struct {
	sum [][]int
}

func NumMatrixConstructor(matrix [][]int) NumMatrix {
	m := NumMatrix{
		sum: make([][]int, len(matrix) + 1),
	}
	for i := 0; i < len(matrix) + 1; i++ {
		m.sum[i] = make([]int, len(matrix[0]) + 1)
	}
	// 记录以原点为0点的矩阵的和
	for i := 1; i < len(m.sum); i++ {
		for j := 1; j < len(m.sum[i]); j++ {
			m.sum[i][j] = m.sum[i - 1][j] + m.sum[i][j - 1] + matrix[i - 1][j - 1] - m.sum[i - 1][j - 1]
		}
	}
	return m
}

// 思路: 以(0, 0)为原点到(row2 + 1, col2 + 1)的矩形 -
//       以(0, 0)为原点到(row1, col2 + 1)的矩形 -
//       以(0, 0)为原点到(row2 + 1, col1)的矩形 +
//       以(0, 0)为原点到(row1, col1)的矩形
func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return n.sum[row2 + 1][col2 + 1] - n.sum[row1][col2 + 1] - n.sum[row2 + 1][col1] + n.sum[row1][col1]
}

// ------------------

type NumMatrix2 struct {
	sum [][]int
}

func NumMatrixConstructor2(matrix [][]int) NumMatrix2 {
	m := NumMatrix2{
		sum: make([][]int, len(matrix) + 1),
	}
	for i := 0; i < len(matrix) + 1; i++ {
		m.sum[i] = make([]int, len(matrix[0]) + 1)
	}
	for i := 1; i < len(m.sum); i++ {
		for j := 1; j < len(m.sum[i]); j++ {
			m.sum[i][j] = m.sum[i][j - 1] + matrix[i - 1][j - 1]
		}
	}
	return m
}

// 思路是计算每一行，两个前缀和相减，每行再相加，虽然能过
// 但是下标贼乱，还容易越界
func (n *NumMatrix2) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := 0
	num := row2 - row1 + 1
	for i := 0; i < num; i++ {
		res += n.sum[row1 + i + 1][col2 + 1] - n.sum[row1 + i + 1][col1]
	}
	return res
}
