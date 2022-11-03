package dfs

func numEnclaves(grid [][]int) int {
	resInNumEnclaves = 0
	for i := 0; i < len(grid); i++ {
		dfsInNumEnclaves(grid, i, 0, true)
		dfsInNumEnclaves(grid, i, len(grid[0]) - 1, true)
	}
	for i := 0; i < len(grid[0]); i++ {
		dfsInNumEnclaves(grid, 0, i, true)
		dfsInNumEnclaves(grid, len(grid) - 1, i, true)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfsInNumEnclaves(grid, i, j, false)
			}
		}
	}
	return resInNumEnclaves
}

var resInNumEnclaves int
func dfsInNumEnclaves(grid [][]int, i, j int, isEdge bool) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	if !isEdge {
		resInNumEnclaves++
	}
	dfsInNumEnclaves(grid, i, j + 1, isEdge)
	dfsInNumEnclaves(grid, i, j - 1, isEdge)
	dfsInNumEnclaves(grid, i + 1, j, isEdge)
	dfsInNumEnclaves(grid, i - 1, j, isEdge)
}
