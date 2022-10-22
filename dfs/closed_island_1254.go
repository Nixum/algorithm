package dfs

func closedIsland(grid [][]int) int {
	// 先把4条边的1都去掉
	for i := 0; i < len(grid); i++ {
		dfsInClosedIsland(grid, i, 0)
		dfsInClosedIsland(grid, i, len(grid[0]) - 1)
	}
	for j := 0; j < len(grid[0]); j++ {
		dfsInClosedIsland(grid, 0, j)
		dfsInClosedIsland(grid, len(grid) - 1, j)
	}
	res := 0
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				res++
				dfsInClosedIsland(grid, i, j)
			}
		}
	}
	return res
}

func dfsInClosedIsland(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfsInClosedIsland(grid, i, j + 1)
	dfsInClosedIsland(grid, i + 1, j)
	dfsInClosedIsland(grid, i - 1, j)
	dfsInClosedIsland(grid, i, j - 1)
}
