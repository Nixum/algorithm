package dfs

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[0]); j++ {
			// 把非子岛屿的先置为排除
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfsInCountSubIslands(grid1, grid2, i, j)
			}
		}
	}
	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[0]); j++ {
			if grid2[i][j] == 1 {
				dfsInCountSubIslands(grid1, grid2, i, j)
				res++
			}
		}
	}
	return res
}

func dfsInCountSubIslands(grid1 [][]int, grid2 [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid2) || j >= len(grid2[0]) {
		return
	}
	if grid2[i][j] == 0 {
		return
	}
	grid2[i][j] = 0
	dfsInCountSubIslands(grid1, grid2, i, j + 1)
	dfsInCountSubIslands(grid1, grid2, i, j - 1)
	dfsInCountSubIslands(grid1, grid2, i + 1, j)
	dfsInCountSubIslands(grid1, grid2, i - 1, j)
}

// 错误的思路，return 会导致有些岛屿没有被置为0
func dfsInCountSubIslands2(grid1 [][]int, grid2 [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid2) || j >= len(grid2[0]) {
		return true
	}
	if grid2[i][j] == 0 {
		return true
	}
	if grid2[i][j] == 1 && grid1[i][j] == 1 {
		grid2[i][j] = 0
		return true
	} else if grid2[i][j] == 1 && grid1[i][j] == 0 {
		grid2[i][j] = 0
		return false
	}
	grid2[i][j] = 0
	return dfsInCountSubIslands2(grid1, grid2, i, j + 1) &&
		dfsInCountSubIslands2(grid1, grid2, i, j - 1) &&
		dfsInCountSubIslands2(grid1, grid2, i + 1, j) &&
		dfsInCountSubIslands2(grid1, grid2, i - 1, j)
}