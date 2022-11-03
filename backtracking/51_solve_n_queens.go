package backtracking

func solveNQueens(n int) [][]string {
	resInSolveNQueues = make([][]string, 0)
	chess := make([][]string, n)
	for i := 0; i < n; i++ {
		chess[i] = make([]string, n)
		for j := 0; j < len(chess[i]); j++ {
			chess[i][j] = "."
		}
	}
	backTrackingInSolveNQueues(chess, n, 0)
	return resInSolveNQueues
}

var resInSolveNQueues [][]string
func backTrackingInSolveNQueues(chess [][]string, n int, x int) {
	if x == n {
		res := make([]string, 0)
		for i := 0; i < n; i++ {
			path := ""
			for j := 0; j < n; j++ {
				path += chess[i][j]
			}
			res = append(res, path)
		}
		resInSolveNQueues = append(resInSolveNQueues, res)
		return
	}
	for j := 0; j < n; j++ {
		if isValidInSolveNQueues(chess, x, j) {
			chess[x][j] = "Q"
			backTrackingInSolveNQueues(chess, n, x + 1)
			chess[x][j] = "."
		}
	}

}

func isValidInSolveNQueues(chess [][]string, row, col int) bool {
	for i := row; i >= 0; i-- {
		if chess[i][col] == "Q" {
			return false
		}
	}
	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if chess[i][j] == "Q" {
			return false
		}
		i--
		j--
	}
	i = row - 1
	j = col + 1
	for i >= 0 && j < len(chess) {
		if chess[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	return true
}