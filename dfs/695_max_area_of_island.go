package dfs

func maxAreaOfIsland(grid [][]int) int {
	maxSize := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfsInMaxAreaOfIsland(grid, i, j)
				if maxSize < size {
					maxSize = size
				}
				size = 0
			}
		}
	}
	return maxSize
}

var size int
func dfsInMaxAreaOfIsland(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	size++
	grid[i][j] = 0
	dfsInMaxAreaOfIsland(grid, i, j + 1)
	dfsInMaxAreaOfIsland(grid, i, j - 1)
	dfsInMaxAreaOfIsland(grid, i + 1, j)
	dfsInMaxAreaOfIsland(grid, i - 1, j)
}
