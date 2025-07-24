package services

/*
Given a 2D grid of '1' (land) and '0' (water), count islands.

Two cells are considered in the same island if they share a vertex or an edge
*/

func CountIsland(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	islandCount := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				islandCount++
				dfs(grid, r, c, rows, cols)
			}
		}
	}
	return islandCount
}

func dfs(grid [][]byte, r, c, rows, cols int) {
	if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c-1, rows, cols)
	dfs(grid, r-1, c, rows, cols)
	dfs(grid, r-1, c+1, rows, cols)
	dfs(grid, r, c-1, rows, cols)
	dfs(grid, r, c+1, rows, cols)
	dfs(grid, r+1, c-1, rows, cols)
	dfs(grid, r+1, c, rows, cols)
	dfs(grid, r+1, c+1, rows, cols)
}
