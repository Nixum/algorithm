package dfs

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfsInNumIslands(grid, i, j)
			}
		}
	}
	return res
}

func dfsInNumIslands(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j >= len(grid[0]) || j < 0 {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfsInNumIslands(grid, i + 1, j)
	dfsInNumIslands(grid, i, j + 1)
	dfsInNumIslands(grid, i - 1, j)
	dfsInNumIslands(grid, i, j - 1)
}