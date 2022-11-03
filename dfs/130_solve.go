package dfs

func solve(board [][]byte)  {
	resInSolve = make([][]int, 0)
	visited := make([][]byte, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]byte, len(board[0]))
		for j := 0; j < len(board[0]); j++ {
			visited[i][j] = board[i][j]
		}
	}
	for i := 0; i < len(visited); i++ {
		dfsInSolve(visited, i, 0, true)
		dfsInSolve(visited, i, len(visited[0]) - 1, true)
	}
	for i := 0; i < len(visited[0]); i++ {
		dfsInSolve(visited, 0, i, true)
		dfsInSolve(visited, len(visited) - 1, i, true)
	}
	for i := 1; i < len(visited) - 1; i++ {
		for j := 1; j < len(visited[0]) - 1; j++ {
			if visited[i][j] == 'O' {
				dfsInSolve(visited, i, j, false)
			}
		}
	}
	for i := 0; i < len(resInSolve); i++ {
		board[resInSolve[i][0]][resInSolve[i][1]] = 'X'
	}
}

var resInSolve [][]int
func dfsInSolve(board [][]byte, i, j int, isEdge bool) {
	if i < 0 || j < 0 ||
		i >= len(board) || j >= len(board[0]) {
		return
	}
	if board[i][j] == 'X' {
		return
	}
	board[i][j] = 'X'
	if !isEdge {
		resInSolve = append(resInSolve, []int{i, j})
	}
	dfsInSolve(board, i, j + 1, isEdge)
	dfsInSolve(board, i, j - 1, isEdge)
	dfsInSolve(board, i + 1, j, isEdge)
	dfsInSolve(board, i - 1, j, isEdge)
}